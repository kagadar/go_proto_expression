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

// Package protoexpr allows for a proto message, and its fields, to be
// declared as an Ident for an API filter string (https://google.aip.dev/160).
package protoexpr

import (
	"fmt"

	"github.com/iancoleman/strcase"
	"go.einride.tech/aip/filtering"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"

	expr "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
	dpb "google.golang.org/protobuf/types/known/durationpb"
	tspb "google.golang.org/protobuf/types/known/timestamppb"

	opb "github.com/kagadar/go_proto_expression/genproto/options"
)

var (
	// WKTs with special handling (we do not need to traverse these messages).
	durationPrefName  protoreflect.FullName = (&dpb.Duration{}).ProtoReflect().Descriptor().FullName()
	timestampPrefName protoreflect.FullName = (&tspb.Timestamp{}).ProtoReflect().Descriptor().FullName()
)

// Map provided proto field to an expr.Type.
// Returns nil if the provided field does not map to a known proto type.
func fieldType(field protoreflect.FieldDescriptor) *expr.Type {
	switch field.Kind() {
	case protoreflect.BoolKind:
		return filtering.TypeBool
	case protoreflect.EnumKind:
		return filtering.TypeEnum(dynamicpb.NewEnumType(field.Enum()))
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Uint32Kind,
		protoreflect.Fixed32Kind, protoreflect.Sfixed32Kind,
		protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed64Kind, protoreflect.Sfixed64Kind:
		return filtering.TypeInt
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return filtering.TypeFloat
	case protoreflect.StringKind, protoreflect.BytesKind:
		return filtering.TypeString
	case protoreflect.MessageKind, protoreflect.GroupKind:
		// Check for WKTs.
		switch name := field.Message().FullName(); name {
		case durationPrefName:
			return filtering.TypeDuration
		case timestampPrefName:
			return filtering.TypeTimestamp
		default:
			return &expr.Type{TypeKind: &expr.Type_MessageType{MessageType: string(name)}}
		}
	}
	return nil
}

func traverseFields(path string, msg protoreflect.MessageDescriptor, msgExpr *expr.Type) []filtering.DeclarationOption {
	options := []filtering.DeclarationOption{
		filtering.DeclareIdent(path, msgExpr),
		filtering.DeclareFunction(filtering.FunctionHas, filtering.NewFunctionOverload(fmt.Sprintf("%s_%s", filtering.FunctionHas, path), filtering.TypeBool, msgExpr, filtering.TypeString)),
	}
	for i := 0; i < msg.Enums().Len(); i++ {
		enum := msg.Enums().Get(i)
		options = append(options, filtering.DeclareEnumIdent(strcase.ToSnake(string(enum.Name())), dynamicpb.NewEnumType(enum)))
	}
	for i := 0; i < msg.Fields().Len(); i++ {
		field := msg.Fields().Get(i)
		if proto.HasExtension(field.Options(), opb.E_Filtering) {
			opts := proto.GetExtension(field.Options(), opb.E_Filtering).(*opb.FieldFilteringOptions)
			// Default value is true; check if set before testing if false.
			if opts.Filterable != nil && !*opts.Filterable {
				continue
			}
		}
		var fType *expr.Type
		switch {
		case field.IsMap():
			if k := fieldType(field.MapKey()); k != nil {
				if v := fieldType(field.MapValue()); v != nil {
					fType = filtering.TypeMap(k, v)
				}
			}
		case field.IsList():
			if t := fieldType(field); fType != nil {
				fType = filtering.TypeList(t)
			}
		default:
			fType = fieldType(field)
		}
		if fType != nil {
			name := fmt.Sprintf("%s.%s", path, field.Name())
			switch field.Kind() {
			case protoreflect.MessageKind, protoreflect.GroupKind:
				options = append(options, traverseFields(name, field.Message(), fType)...)
			default:
				options = append(options, filtering.DeclareIdent(name, fType))
			}
		}
	}
	return options
}

// Iterate through all fields of a supplied protobuf Message to create Ident
// Declarations.
// This will produce Ident Declarations for the protobuf Message itself, as well
// as an overload for the `Has` function to check for field presence.
//
// Each field will have an Ident produced with an appropriate expr.Type.
// All enums declared within the message will have their values declared as
// constants.
// All fiels of message type will processed recursively, with their Indent name
// set to the relative path from the root message.
// google.protobuf.Timestamp and google.protobuf.Duration are well known types,
// and will return an appropriate WellKnown type rather than a Message type.
func Declare(msg protoreflect.MessageDescriptor) []filtering.DeclarationOption {
	return traverseFields(strcase.ToSnake(string(msg.Name())), msg, &expr.Type{TypeKind: &expr.Type_MessageType{MessageType: string(msg.FullName())}})
}
