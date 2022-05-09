// Copyright 2022 The Go Protobuf Expression Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package protoexpr

import (
	"testing"

	"go.einride.tech/aip/filtering"

	"github.com/kagadar/go_proto_expression/protoexpr/test"
)

type req struct {
	filter string
}

func (r req) GetFilter() string {
	return r.filter
}

func TestDeclare(t *testing.T) {
	decls, err := filtering.NewDeclarations(append(append([]filtering.DeclarationOption{}, filtering.DeclareStandardFunctions()), Declare((&test.TestFiltering{}).ProtoReflect().Descriptor())...)...)
	if err != nil {
		t.Fatalf(`Declare() err = %v, want <nil>`, err)
	}
	if _, err := filtering.ParseFilter(req{
		`
			test_filtering:filterable_submessage AND
			test_filtering.filterable_submessage.filterable_primitive = 1 AND
			test_filtering:default_submessage AND
			test_filtering.default_submessage.filterable_primitive = 2 AND
			test_filtering.filterable_primitive = "test" AND
			test_filtering.default_float = 1.0 AND
			test_filtering.default_bool AND
			test_filtering.default_enum = VALUE_1
		`,
	}, decls); err != nil {
		t.Fatalf(`filtering.ParseFilter(filterable) err = %v, want <nil>`, err)
	}
	if _, err := filtering.ParseFilter(req{`test_filtering.filterable_submessage.unfilterable_primitive = "test"`}, decls); err == nil {
		t.Fatal(`filtering.ParseFilter(unfilterable message field) err = <nil>, want error`)
	}
	if _, err := filtering.ParseFilter(req{`test_filtering.unfilterable_submessage.filterable_primitive = 1`}, decls); err == nil {
		t.Fatal(`filtering.ParseFilter(unfilterable message) err = <nil>, want error`)
	}
	if _, err := filtering.ParseFilter(req{`test_filtering.unfilterable_primitive = 1`}, decls); err == nil {
		t.Fatal(`filtering.ParseFilter(unfilterable field) err = <nil>, want error`)
	}
}
