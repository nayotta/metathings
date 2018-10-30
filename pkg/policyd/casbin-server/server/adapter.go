// Copyright 2018 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"errors"

	"github.com/casbin/casbin/persist"
	"github.com/casbin/casbin/persist/file-adapter"
	gormadapter "github.com/casbin/gorm-adapter"
	// _ "github.com/jinzhu/gorm/dialects/mssql"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	// _ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	pb "github.com/nayotta/metathings/pkg/proto/policyd"
)

var errDriverName = errors.New("invalid DriverName")

const (
	FILE = "file"
	GORM = "gorm"
)

func newAdapter(in *pb.NewAdapterRequest) (persist.Adapter, error) {
	var a persist.Adapter

	switch in.DriverName {
	case FILE:
		a = fileadapter.NewAdapter(in.ConnectString)
	case GORM:
		a = gormadapter.NewAdapter(in.AdapterName, in.ConnectString)
	default:
		return nil, errDriverName
	}

	return a, nil
}
