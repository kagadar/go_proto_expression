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

	"go.einride.tech/aip/filtering"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"

	opb "github.com/kagadar/go_proto_expression/genproto/options"
)

// AIP-132 & AIP-160 compliant List Request.
type ListRequest interface {
	proto.Message
	GetParent() string
	GetPageSize() int32
	GetPageToken() string
	GetFilter() string
}

// Client provides a concrete implementation of the Transpiler for a specific storage client.
type Client[T proto.Message] interface {
	Transpile(ctx context.Context, factory func() T, parent, collection, pageToken string, pageSize int32, filter filtering.Filter) (children []T, nextPageToken string, err error)
}

type Transpiler[T proto.Message] interface {
	Transpile(context.Context, ListRequest) (children []T, nextPageToken string, err error)
}

type transpiler[T proto.Message] struct {
	client          Client[T]
	collection      string
	decls           *filtering.Declarations
	defaultPageSize int32
	maxPageSize     int32
	emptyMessage    proto.Message
}

func (t transpiler[T]) Transpile(ctx context.Context, req ListRequest) ([]T, string, error) {
	pageSize := req.GetPageSize()
	switch {
	case pageSize < 0:
		return nil, "", status.Errorf(codes.InvalidArgument, "page size cannot be negative")
	case pageSize == 0:
		pageSize = t.defaultPageSize
	case req.GetPageSize() > t.maxPageSize:
		pageSize = t.maxPageSize
	}
	filter, err := filtering.ParseFilter(req, t.decls)
	if err != nil {
		return nil, "", err
	}
	return t.client.Transpile(ctx, func() T { return proto.Clone(t.emptyMessage).(T) }, req.GetParent(), t.collection, req.GetPageToken(), pageSize, filter)
}

// Creates a new transpiler for requests to the specified List method.
// This should be used by the concrete Transpiler implementation.
func New[T proto.Message](client Client[T], mtd protoreflect.MethodDescriptor, msg T) (Transpiler[T], error) {
	// Check fields on request and response.
	if fields := mtd.Input().Fields(); fields.ByName("parent") == nil || fields.ByName("page_size") == nil || fields.ByName("page_token") == nil || fields.ByName("filter") == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request for method %q is not compliant with AIP-132 and AIP-160", mtd.Name())
	}
	if mtd.Output().Fields().ByName("next_page_token") == nil {
		return nil, status.Errorf(codes.InvalidArgument, "response for method %q is not compliant with AIP-132", mtd.Name())
	}
	t := transpiler[T]{
		client:          client,
		defaultPageSize: 10,
		maxPageSize:     100,
		emptyMessage:    proto.Clone(msg),
	}
	proto.Reset(t.emptyMessage)
	if proto.HasExtension(mtd.Options(), opb.E_Pagination) {
		options := proto.GetExtension(mtd.Options(), opb.E_Pagination).(*opb.MethodPaginationOptions)
		if options.DefaultPageSize != nil {
			t.defaultPageSize = options.GetDefaultPageSize()
		}
		if options.MaxPageSize != nil {
			t.maxPageSize = options.GetMaxPageSize()
		}
	}
	var collectionFieldNum int32 = 1
	if proto.HasExtension(mtd.Output().Options(), opb.E_Collection) {
		options := proto.GetExtension(mtd.Output().Options(), opb.E_Collection).(*opb.MessageCollectionOptions)
		if options.CollectionFieldNumber != nil {
			collectionFieldNum = options.GetCollectionFieldNumber()
		}
	}
	collectionField := mtd.Output().Fields().ByNumber(protowire.Number(collectionFieldNum))
	if collectionField == nil || !collectionField.IsList() || collectionField.Kind() != protoreflect.MessageKind || collectionField.Message() != t.emptyMessage.ProtoReflect().Descriptor() {
		return nil, status.Errorf(codes.InvalidArgument, "unable to determine collection field for %s", mtd.Output().FullName())
	}
	t.collection = string(collectionField.Name())
	decls, err := filtering.NewDeclarations(append(append([]filtering.DeclarationOption{}, filtering.DeclareStandardFunctions()), Declare(collectionField.Message())...)...)
	if err != nil {
		return nil, err
	}
	t.decls = decls
	return t, nil
}
