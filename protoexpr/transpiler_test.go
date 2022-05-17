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
	"context"
	"testing"

	"google.golang.org/protobuf/proto"

	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"

	"github.com/kagadar/go_proto_expression/protoexpr/test"
)

type transpilerClient[T proto.Message] struct{}

func (transpilerClient[T]) Transpile(ctx context.Context, generator func() T, parent, collection, pageToken string, pageSize int32, filter *expr.CheckedExpr) (children []T, nextPageToken string, err error) {
	return
}

func TestTranspiler(t *testing.T) {
	transpiler, err := New[*test.TestFiltering](transpilerClient[*test.TestFiltering]{}, test.File_protoexpr_protoexpr_test_proto.Services().ByName("TestService").Methods().ByName("ListTest"), &test.TestFiltering{})
	if err != nil {
		t.Fatalf("New() err = %v, want <nil>", err)
	}
	if _, _, err := transpiler.Transpile(context.Background(), &test.ListTestRequest{}); err != nil {
		t.Fatalf("transpiler.Transpile() err = %v, want <nil>", err)
	}
}
