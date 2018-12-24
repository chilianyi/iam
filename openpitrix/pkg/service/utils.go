// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package service

import (
	"reflect"
	"strings"

	"github.com/chai2010/spacestring"
	"github.com/fatih/structs"
)

func pkgGetTableFiledNamesAndValues(v interface{}) (names []string, values []interface{}) {
	s := structs.New(v)
	for _, f := range s.Fields() {
		if !f.IsExported() || f.IsZero() {
			continue
		}
		if strings.HasPrefix(f.Name(), "XXX_") || f.Tag("json") == "-" {
			continue
		}

		var (
			db_field_name  = f.Name()
			db_field_value = f.Value()
		)

		if idx := strings.Index(f.Tag("json"), ","); idx > 0 {
			db_field_name = f.Tag("json")[:idx]
		}

		if f.Kind() == reflect.String {
			if spacestring.IsSpace(f.Value().(string)) {
				db_field_value = "" // clear field
			}
		}

		// TODO(chai): support pb.Timestamp
		// TODO(chai): keep empty value for insert into

		names = append(names, db_field_name)
		values = append(values, db_field_value)
	}
	return
}
