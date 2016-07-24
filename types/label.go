// Copyright 2013-2016 lessgo Author, All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	labelPatName        = regexp.MustCompile("^[a-z]{1}[a-z0-9-._/]{2,100}$")
	labelErrNameLength  = errors.New("Length of the label name must be between 5 and 100")
	labelErrNameInvalid = errors.New("Invalid Label Name")
)

// Labels are name value pairs that may be used to scope and select individual items.
type Labels []Label

// Label is a name value . It implements Labels
type Label struct {
	Name  string `json:"name"`
	Value string `json:"value,omitempty"`
}

// Set create or update the label entry for "name" to "value".
func (ls *Labels) Set(name string, value interface{}) error {

	if len(name) < 4 || len(name) > 100 {
		return labelErrNameLength
	}

	if !labelPatName.MatchString(name) {
		return labelErrNameInvalid
	}

	svalue := fmt.Sprintf("%v", value)

	for i, prev := range *ls {

		if prev.Name == name {

			if prev.Value != svalue {
				(*ls)[i].Value = svalue
			}

			return nil
		}
	}

	(*ls) = append(*ls, Label{
		Name:  name,
		Value: svalue,
	})

	return nil
}

// Get fetch the label entry "value" (if any) for "name".
func (ls Labels) Get(name string) (Bytex, bool) {

	for _, prev := range ls {

		if prev.Name == name {
			return Bytex(prev.Value), true
		}
	}

	return Bytex(""), false
}

// Del remove the label entry (if any) for "name".
func (ls *Labels) Del(name string) {

	for i, prev := range *ls {

		if prev.Name == name {
			(*ls) = append((*ls)[:i], (*ls)[i+1:]...)
			break
		}
	}
}

func (ls *Labels) Equal(items Labels) bool {

	if len(*ls) != len(items) {
		return false
	}

	for _, v := range *ls {

		hit := false

		for _, v2 := range items {

			if v.Name != v2.Name {
				continue
			}

			if v.Value != v2.Value {
				return false
			}

			hit = true
			break
		}

		if !hit {
			return false
		}
	}

	return true
}
