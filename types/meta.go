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

// Common string formats
// ---------------------
// Many fields in this API have formatting requirements.  The commonly used
// formats are defined here.
//
// C_IDENTIFIER:  This is a string that conforms to the definition of an "identifier"
//     in the C language.  This is captured by the following regex:
//         [A-Za-z_][A-Za-z0-9_]*
//     This defines the format, but not the length restriction, which should be
//     specified at the definition of any field of this type.
//
// DNS_LABEL:  This is a string, no more than 63 characters long, that conforms
//     to the definition of a "label" in RFCs 1035 and 1123.  This is captured
//     by the following regex:
//         [a-z0-9]([-a-z0-9]*[a-z0-9])?
//
// DNS_SUBDOMAIN:  This is a string, no more than 253 characters long, that conforms
//      to the definition of a "subdomain" in RFCs 1035 and 1123.  This is captured
//      by the following regex:
//         [a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*
//     or more simply:
//         DNS_LABEL(\.DNS_LABEL)*

// TypeMeta describes an individual object in an API response or request
// with strings representing the type of the object and its API schema version.
// Structures that are versioned or persisted should inline TypeMeta.
type TypeMeta struct {
	// APIVersion defines the versioned schema of this representation of an object.
	// Servers should convert recognized schemas to the latest internal value, and
	// may reject unrecognized values.
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind is a string value representing the REST resource this object represents.
	// Servers may infer this from the endpoint the client submits requests to.
	Kind string `json:"kind,omitempty"`

	// ErrorMeta provides more information about an api failure. If this value is
	// nil there is no error available
	Error *ErrorMeta `json:"error,omitempty"`
}

func NewTypeErrorMeta(code, message string) TypeMeta {
	return TypeMeta{
		Error: &ErrorMeta{
			Code:    code,
			Message: message,
		},
	}
}

// ErrorMeta provides more information about an api failure.
type ErrorMeta struct {
	// A machine-readable description of the type of the error. If this value is
	// empty there is no information available.
	Code string `json:"code,omitempty"`

	// A human-readable description of the error message.
	Message string `json:"message,omitempty"`
}

func NewErrorMeta(code, message string) *ErrorMeta {
	return &ErrorMeta{
		Code:    code,
		Message: message,
	}
}

// ListMeta describes metadata that synthetic resources must have, including lists and
// various status objects. A resource may have only one of {ObjectMeta, ListMeta}.
type ListMeta struct {
	// SelfLink is a URL representing this object.
	SelfLink string `json:"selfLink,omitempty"`

	// An opaque value that represents the version of this response for use with optimistic
	// concurrency and change monitoring endpoints.  Clients must treat these values as opaque
	// and values may only be valid for a particular resource or set of resources. Only servers
	// will generate resource versions.
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// The number of query results available for the current query
	TotalResults uint64 `json:"totalResults,omitempty"`

	// The index of the first query result in the current set of query results
	StartIndex uint64 `json:"startIndex,omitempty"`

	// The number of query results returned per list
	ItemsPerList uint64 `json:"itemsPerList,omitempty"`
}

// ObjectMeta is metadata that all persisted resources must have, which includes all objects
// users must create. A resource may have only one of {ObjectMeta, ListMeta}.
type ObjectMeta struct {
	// ID is the unique in time and space value for this object. It is typically generated by
	// the server on successful creation of a resource and is not allowed to change on PUT
	// operations.
	ID string `json:"id,omitempty"`

	// Name is unique within a namespace.  Name is required when creating resources, although
	// some resources may allow a client to request the generation of an appropriate name
	// automatically. Name is primarily intended for creation idempotence and configuration
	// definition.
	Name string `json:"name,omitempty"`

	// UserID is the unique in time and space value for a user.
	UserID string `json:"userID,omitempty"`

	// An opaque value that represents the version of this resource. May be used for optimistic
	// concurrency, change detection, and the watch operation on a resource or set of resources.
	// Clients must treat these values as opaque and values may only be valid for a particular
	// resource or set of resources. Only servers will generate resource versions.
	ResourceVersion string `json:"resourceVersion,omitempty"`

	// Labels are key value pairs that may be used to scope and select individual resources.
	Labels map[string]string `json:"labels,omitempty"`

	// Created or Updated is a timestamp representing the server time when this object was created
	// or updated. It is not guaranteed to be set in happens-before order across separate operations.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	Created string `json:"created,omitempty"`
	Updated string `json:"updated,omitempty"`
}

// ObjectMeta is metadata that all persisted resources must have, which includes all objects
// users must create. A resource may have only one of {ObjectMeta, ListMeta}.
type InnerObjectMeta struct {
	// ID is the unique in time and space value for this object. It is typically generated by
	// the server on successful creation of a resource and is not allowed to change on PUT
	// operations.
	ID string `json:"id,omitempty"`

	// Name is unique within a namespace.  Name is required when creating resources, although
	// some resources may allow a client to request the generation of an appropriate name
	// automatically. Name is primarily intended for creation idempotence and configuration
	// definition.
	Name string `json:"name,omitempty"`

	// User defines the username or userid.
	User string `json:"user,omitempty"`

	// An opaque value that represents the version of this resource. May be used for optimistic
	// concurrency, change detection, and the watch operation on a resource or set of resources.
	// Clients must treat these values as opaque and values may only be valid for a particular
	// resource or set of resources. Only servers will generate resource versions.
	Version string `json:"version,omitempty"`

	// Created or Updated is a timestamp representing the server time when this object was created
	// or updated. It is not guaranteed to be set in happens-before order across separate operations.
	// Clients may not set this value. It is represented in RFC3339 form and is in UTC.
	Created MetaTime `json:"created,omitempty"`
	Updated MetaTime `json:"updated,omitempty"`

	// Human readable description of this object.
	Title string `json:"title,omitempty"`
}