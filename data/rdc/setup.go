package rdc

import (
	"./mysql"
	"./setup"
	"./sqlite3"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func (cn *Conn) FieldType(t string) string {

	switch cn.driver {
	case "mysql":
		return mysql.FieldTypes[t]
	case "sqlite3":
		return sqlite3.FieldTypes[t]
	}

	return ""
}

//
func (cn *Conn) Setup(dsname string, ds setup.DataSet) error {

	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS less_dataset_version (
        id %s,
        version %s,
        action %s,
        created %s);`,
		cn.FieldType("auto"),
		cn.FieldType("int32"),
		fmt.Sprintf(cn.FieldType("string"), 30),
		cn.FieldType("datetime"))
	if _, err := cn.ExecRaw(sql); err != nil {
		return err
	}

	dsVerPrev := 0
	q := NewQuerySet().From("less_dataset_version").Order("version desc").Limit(1)
	if rs, err := cn.Query(q); err != nil {
		return err
	} else if len(rs) == 1 {
		dsVerPrev = int(reflect.ValueOf(rs[0]["version"]).Int())
		if dsVerPrev >= int(ds.Version) {
			return nil
		}
	}

	//var err error
	for _, v := range ds.Tables {

		update := false

		// TODO driver/sqlite*
		// cid name type notnull dflt_value pk
		fscur, err := cn.QueryRaw("PRAGMA table_info(" + v.Name + ")")
		if err != nil {
			return err
		}
		if len(fscur) == 0 {
			update = true
		}

		if !update {

			// New Field Append
			for _, v2 := range v.Fields {

				isnew := true

				for _, v3 := range fscur {

					if v2.Name == v3["name"] {

						isnew = false

						if strings.ToLower(v3["type"].(string)) != v2.Type {
							update = true
						}
					}
				}

				if isnew {
					update = true
				}

				if update {
					break
				}
			}
		}

		if !update {
			for _, v2 := range fscur {

				delete := true

				for _, v3 := range v.Fields {
					if v3.Name == v2["name"] {
						delete = false
					}
				}

				if delete {
					update = true
					break
				}
			}
		}

		if !update {
			continue
		}

		//
		backup := ""
		if len(fscur) > 0 {
			backup = v.Name + "_" + strconv.Itoa(dsVerPrev) + "_" + time.Now().Format("20060102_150405")
			// TODO driver/sqlite*
			sql := fmt.Sprintf("ALTER TABLE %s RENAME TO %s", v.Name, backup)
			if _, err = cn.db.Exec(sql); err != nil {
				return err
			}
		}

		//
		fs := []string{}
		for _, v2 := range v.Fields {
			if fstr := cn.FieldType(v2.Type); len(fstr) > 0 {

				switch v2.Type {
				case "string", "pk-string":
					if v2.Len < 1 || v2.Len > 256 {
						v2.Len = 50
					}
					fstr = fmt.Sprintf(fstr, v2.Len)
				}

				fs = append(fs, "`"+v2.Name+"` "+fstr)
			}
		}

		sql := fmt.Sprintf("CREATE TABLE %s (%s)", v.Name, strings.Join(fs, ","))
		if _, err = cn.db.Exec(sql); err != nil {
			return err
		}
		for _, v2 := range v.Fields {

			action := ""
			if v2.Idx == setup.FieldIndexIndex {
				action += "CREATE INDEX "
			} else if v2.Idx == setup.FieldIndexUnique {
				action += "CREATE UNIQUE INDEX "
			}

			if len(action) > 0 {
				sql = fmt.Sprintf("%s %s_%s_idx ON %s (%s)",
					action, v.Name, v2.Name, v.Name, v2.Name)
				if _, err = cn.db.Exec(sql); err != nil {
					//return err
				}
			}
		}

		//
		if len(fscur) > 0 {

			fs := []string{}
			for _, v2 := range v.Fields {
				for _, v3 := range fscur {
					if v2.Name == v3["name"] {
						fs = append(fs, "`"+v2.Name+"`")
					}
				}
			}
			sql := fmt.Sprintf("INSERT INTO `%s` (%s) SELECT %s FROM %s",
				v.Name, strings.Join(fs, ","), strings.Join(fs, ","), backup)
			//fmt.Println("sql", sql)
			if _, err = cn.db.Exec(sql); err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	item := map[string]interface{}{
		"version": ds.Version,
		"action":  "update",
		"created": time.Now().Format(timeFormatDateTime),
	}
	_, _ = cn.Insert("less_dataset_version", item)

	return nil
}
