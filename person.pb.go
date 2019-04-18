// Code generated by protoc-gen-go. DO NOT EDIT.
// source: person.proto

/*
Package main is a generated protocol buffer package.

It is generated from these files:
	person.proto

It has these top-level messages:
	BMC
	Person
*/
package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type BMC struct {
	Height int32 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
	Weight int32 `protobuf:"varint,2,opt,name=weight" json:"weight,omitempty"`
}

func (m *BMC) Reset()                    { *m = BMC{} }
func (m *BMC) String() string            { return proto.CompactTextString(m) }
func (*BMC) ProtoMessage()               {}
func (*BMC) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BMC) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BMC) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type Person struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,2,opt,name=age" json:"age,omitempty"`
	Bmc  *BMC   `protobuf:"bytes,3,opt,name=bmc" json:"bmc,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Person) GetBmc() *BMC {
	if m != nil {
		return m.Bmc
	}
	return nil
}

func init() {
	proto.RegisterType((*BMC)(nil), "main.BMC")
	proto.RegisterType((*Person)(nil), "main.Person")
}

func init() { proto.RegisterFile("person.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x2d, 0x2a,
	0xce, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x4d, 0xcc, 0xcc, 0x53, 0x32,
	0xe5, 0x62, 0x76, 0xf2, 0x75, 0x16, 0x12, 0xe3, 0x62, 0xcb, 0x48, 0xcd, 0x4c, 0xcf, 0x28, 0x91,
	0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf2, 0x40, 0xe2, 0xe5, 0x10, 0x71, 0x26, 0x88, 0x38,
	0x84, 0xa7, 0xe4, 0xcd, 0xc5, 0x16, 0x00, 0x36, 0x4c, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37,
	0x15, 0xac, 0x8f, 0x33, 0x08, 0xcc, 0x16, 0x12, 0xe0, 0x62, 0x4e, 0x4c, 0x4f, 0x85, 0x6a, 0x01,
	0x31, 0x85, 0xa4, 0xb9, 0x98, 0x93, 0x72, 0x93, 0x25, 0x98, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x38,
	0xf5, 0x40, 0x56, 0xeb, 0x39, 0xf9, 0x3a, 0x07, 0x81, 0x44, 0x93, 0xd8, 0xc0, 0x0e, 0x32, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x7d, 0xd7, 0x4e, 0xa0, 0x00, 0x00, 0x00,
}
