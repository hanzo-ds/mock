// Licensed to Datastore, Inc. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Datastore, Inc. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// This file is a copy of the original file from the datastore-go package.
// The original file is located at:
// https://github.com/ClickHouse/datastore-go/blob/226a902d120aa46e3883fbf6a5a2667dfb9e90d2/datastore_rows_column_type.go

package mock

import (
	"reflect"

	"github.com/hanzo-ds/go/lib/driver"
)

type columnType struct {
	name     string
	dsType   string
	nullable bool
	scanType reflect.Type
}

func NewColumnType(name string, dsType string, nullable bool, scanType reflect.Type) driver.ColumnType {
	return columnType{
		name:     name,
		dsType:   dsType,
		nullable: nullable,
		scanType: scanType,
	}
}

func (c columnType) Name() string {
	return c.name
}

func (c columnType) Nullable() bool {
	return c.nullable
}

func (c columnType) ScanType() reflect.Type {
	return c.scanType
}

func (c columnType) DatabaseTypeName() string {
	return c.dsType
}
