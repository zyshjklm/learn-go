// Code generated by protoc-gen-go. DO NOT EDIT.
// source: myproto/addressbook.proto

/*
Package myproto is a generated protocol buffer package.

It is generated from these files:
	myproto/addressbook.proto

It has these top-level messages:
	PhoneNumber
	Person
	AddressBook
*/
package myproto

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

type PhoneType int32

const (
	PhoneType_MOBILE PhoneType = 0
	PhoneType_HOME   PhoneType = 1
	PhoneType_WORK   PhoneType = 2
)

var PhoneType_name = map[int32]string{
	0: "MOBILE",
	1: "HOME",
	2: "WORK",
}
var PhoneType_value = map[string]int32{
	"MOBILE": 0,
	"HOME":   1,
	"WORK":   2,
}

func (x PhoneType) String() string {
	return proto.EnumName(PhoneType_name, int32(x))
}
func (PhoneType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PhoneNumber struct {
	Number string    `protobuf:"bytes,1,opt,name=number" json:"number,omitempty"`
	Type   PhoneType `protobuf:"varint,2,opt,name=type,enum=myproto.PhoneType" json:"type,omitempty"`
}

func (m *PhoneNumber) Reset()                    { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string            { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()               {}
func (*PhoneNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func (m *PhoneNumber) GetType() PhoneType {
	if m != nil {
		return m.Type
	}
	return PhoneType_MOBILE
}

type Person struct {
	Id     int32          `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name   string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Email  string         `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Phones []*PhoneNumber `protobuf:"bytes,4,rep,name=phones" json:"phones,omitempty"`
}

func (m *Person) Reset()                    { *m = Person{} }
func (m *Person) String() string            { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()               {}
func (*Person) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Person) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Person) GetPhones() []*PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type AddressBook struct {
	People []*Person `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *AddressBook) Reset()                    { *m = AddressBook{} }
func (m *AddressBook) String() string            { return proto.CompactTextString(m) }
func (*AddressBook) ProtoMessage()               {}
func (*AddressBook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AddressBook) GetPeople() []*Person {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*PhoneNumber)(nil), "myproto.PhoneNumber")
	proto.RegisterType((*Person)(nil), "myproto.Person")
	proto.RegisterType((*AddressBook)(nil), "myproto.AddressBook")
	proto.RegisterEnum("myproto.PhoneType", PhoneType_name, PhoneType_value)
}

func init() { proto.RegisterFile("myproto/addressbook.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0x34, 0x5d, 0xcd, 0x04, 0x6a, 0x18, 0x8a, 0xac, 0xb7, 0x90, 0x83, 0x06, 0x95,
	0x08, 0x15, 0xbc, 0x5b, 0x28, 0x28, 0x1a, 0x53, 0x16, 0xc1, 0x73, 0x42, 0x16, 0x0c, 0x6d, 0x32,
	0x4b, 0x52, 0x0f, 0xf9, 0xf7, 0xd2, 0xc9, 0xa2, 0x78, 0x7b, 0x33, 0x6f, 0xdf, 0xb7, 0x6f, 0xe0,
	0xb2, 0x1d, 0x6d, 0x4f, 0x07, 0xba, 0x2f, 0xeb, 0xba, 0x37, 0xc3, 0x50, 0x11, 0xed, 0x32, 0xde,
	0xe0, 0xa9, 0xb3, 0x92, 0x1c, 0xc2, 0xed, 0x17, 0x75, 0xe6, 0xfd, 0xbb, 0xad, 0x4c, 0x8f, 0x17,
	0x20, 0x3b, 0x56, 0x4a, 0xc4, 0x22, 0x0d, 0xb4, 0x9b, 0xf0, 0x0a, 0xfc, 0xc3, 0x68, 0x8d, 0xf2,
	0x62, 0x91, 0x2e, 0x56, 0x98, 0xb9, 0x78, 0xc6, 0xd9, 0x8f, 0xd1, 0x1a, 0xcd, 0x7e, 0x62, 0x41,
	0x6e, 0x4d, 0x3f, 0x50, 0x87, 0x0b, 0xf0, 0x9a, 0x9a, 0x29, 0x73, 0xed, 0x35, 0x35, 0x22, 0xf8,
	0x5d, 0xd9, 0x4e, 0x84, 0x40, 0xb3, 0xc6, 0x25, 0xcc, 0x4d, 0x5b, 0x36, 0x7b, 0x35, 0xe3, 0xe5,
	0x34, 0xe0, 0x1d, 0x48, 0x7b, 0xc4, 0x0e, 0xca, 0x8f, 0x67, 0x69, 0xb8, 0x5a, 0xfe, 0xff, 0x6d,
	0x6a, 0xaa, 0xdd, 0x9b, 0xe4, 0x11, 0xc2, 0xa7, 0xe9, 0xbc, 0x35, 0xd1, 0x0e, 0xaf, 0x41, 0x5a,
	0x43, 0x76, 0x6f, 0x94, 0xe0, 0xf0, 0xf9, 0x5f, 0x98, 0x7b, 0x69, 0x67, 0xdf, 0xdc, 0x42, 0xf0,
	0x5b, 0x1e, 0x01, 0x64, 0x5e, 0xac, 0x5f, 0xde, 0x36, 0xd1, 0x09, 0x9e, 0x81, 0xff, 0x5c, 0xe4,
	0x9b, 0x48, 0x1c, 0xd5, 0x67, 0xa1, 0x5f, 0x23, 0xaf, 0x92, 0x8c, 0x78, 0xf8, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xf9, 0x9d, 0x18, 0xb0, 0x52, 0x01, 0x00, 0x00,
}
