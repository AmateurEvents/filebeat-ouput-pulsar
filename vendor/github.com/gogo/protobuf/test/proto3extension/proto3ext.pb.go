// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto3ext.proto

package proto3extension

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

var E_Primary = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         51234,
	Name:          "proto3extension.primary",
	Tag:           "varint,51234,opt,name=primary",
	Filename:      "proto3ext.proto",
}

var E_Index = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         51235,
	Name:          "proto3extension.index",
	Tag:           "varint,51235,opt,name=index",
	Filename:      "proto3ext.proto",
}

func init() {
	proto.RegisterExtension(E_Primary)
	proto.RegisterExtension(E_Index)
}

func init() { proto.RegisterFile("proto3ext.proto", fileDescriptor_a399cd792699f40f) }

var fileDescriptor_a399cd792699f40f = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0x37, 0x4e, 0xad, 0x28, 0xd1, 0x03, 0xb3, 0x84, 0x10, 0x02, 0xa9, 0x79, 0xc5, 0x99, 0xf9,
	0x79, 0x52, 0x0a, 0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x60, 0xf1, 0xa4, 0xd2, 0x34, 0xfd,
	0x94, 0xd4, 0xe2, 0xe4, 0xa2, 0xcc, 0x82, 0x92, 0xfc, 0x22, 0x88, 0x16, 0x2b, 0x4b, 0x2e, 0xf6,
	0x82, 0xa2, 0xcc, 0xdc, 0xc4, 0xa2, 0x4a, 0x21, 0x59, 0x3d, 0x88, 0x6a, 0x3d, 0x98, 0x6a, 0x3d,
	0xb7, 0xcc, 0xd4, 0x9c, 0x14, 0xff, 0x82, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x89, 0x45, 0x13, 0x98,
	0x15, 0x18, 0x35, 0x38, 0x82, 0x60, 0xea, 0xad, 0x4c, 0xb9, 0x58, 0x33, 0xf3, 0x52, 0x52, 0x2b,
	0x08, 0x69, 0x5c, 0x0c, 0xd5, 0x08, 0x51, 0x9d, 0xc4, 0x06, 0x71, 0x24, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xff, 0xd4, 0x32, 0x01, 0xbe, 0x00, 0x00, 0x00,
}
