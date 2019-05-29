// Code generated by protoc-gen-go. DO NOT EDIT.
// source: student.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Student struct {
	Name   string         `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Age    uint32         `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
	Emails []string       `protobuf:"bytes,3,rep,name=Emails,proto3" json:"Emails,omitempty"`
	Phones []*PhoneNumber `protobuf:"bytes,4,rep,name=Phones,proto3" json:"Phones,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*Student_School
	//	*Student_Socre
	Data                 isStudent_Data `protobuf_oneof:"Data"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{0}
}

func (m *Student) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Student.Unmarshal(m, b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Student.Marshal(b, m, deterministic)
}
func (m *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(m, src)
}
func (m *Student) XXX_Size() int {
	return xxx_messageInfo_Student.Size(m)
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Student) GetAge() uint32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Student) GetEmails() []string {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *Student) GetPhones() []*PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type isStudent_Data interface {
	isStudent_Data()
}

type Student_School struct {
	School string `protobuf:"bytes,5,opt,name=School,proto3,oneof"`
}

type Student_Socre struct {
	Socre uint32 `protobuf:"varint,6,opt,name=Socre,proto3,oneof"`
}

func (*Student_School) isStudent_Data() {}

func (*Student_Socre) isStudent_Data() {}

func (m *Student) GetData() isStudent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Student) GetSchool() string {
	if x, ok := m.GetData().(*Student_School); ok {
		return x.School
	}
	return ""
}

func (m *Student) GetSocre() uint32 {
	if x, ok := m.GetData().(*Student_Socre); ok {
		return x.Socre
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Student) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Student_School)(nil),
		(*Student_Socre)(nil),
	}
}

type PhoneNumber struct {
	Number               string   `protobuf:"bytes,1,opt,name=Number,proto3" json:"Number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhoneNumber) Reset()         { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()    {}
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{1}
}

func (m *PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneNumber.Unmarshal(m, b)
}
func (m *PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneNumber.Merge(m, src)
}
func (m *PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_PhoneNumber.Size(m)
}
func (m *PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneNumber proto.InternalMessageInfo

func (m *PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func init() {
	proto.RegisterType((*Student)(nil), "pb.Student")
	proto.RegisterType((*PhoneNumber)(nil), "pb.PhoneNumber")
}

func init() { proto.RegisterFile("student.proto", fileDescriptor_94a1c1b032ad0c00) }

var fileDescriptor_94a1c1b032ad0c00 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xb1, 0x6a, 0x86, 0x30,
	0x14, 0x85, 0xff, 0x18, 0x4d, 0xf1, 0x8a, 0xb4, 0xdc, 0x41, 0x32, 0x06, 0xa1, 0x34, 0x93, 0x43,
	0xfb, 0x04, 0x2d, 0x2d, 0x38, 0x49, 0x89, 0x4f, 0x90, 0xd8, 0x50, 0x0b, 0x6a, 0x44, 0xe3, 0x4b,
	0xf5, 0x29, 0x4b, 0x34, 0x43, 0xb7, 0xef, 0x3b, 0x81, 0x9c, 0x73, 0xa1, 0xdc, 0xfd, 0xf1, 0x65,
	0x17, 0xdf, 0xac, 0x9b, 0xf3, 0x0e, 0x93, 0xd5, 0xd4, 0xbf, 0x04, 0xee, 0xfa, 0x2b, 0x45, 0x84,
	0xb4, 0xd3, 0xb3, 0xe5, 0x44, 0x10, 0x99, 0xab, 0x93, 0xf1, 0x01, 0xe8, 0xeb, 0xb7, 0xe5, 0x89,
	0x20, 0xb2, 0x54, 0x01, 0xb1, 0x02, 0xf6, 0x31, 0xeb, 0x9f, 0x69, 0xe7, 0x54, 0x50, 0x99, 0xab,
	0x68, 0xf8, 0x04, 0xec, 0x73, 0x74, 0x8b, 0xdd, 0x79, 0x2a, 0xa8, 0x2c, 0x9e, 0xef, 0x9b, 0xd5,
	0x34, 0x67, 0xd2, 0x1d, 0xb3, 0xb1, 0x9b, 0x8a, 0xcf, 0xc8, 0x81, 0xf5, 0xc3, 0xe8, 0xdc, 0xc4,
	0xb3, 0x50, 0xd4, 0xde, 0x54, 0x74, 0xac, 0x20, 0xeb, 0xdd, 0xb0, 0x59, 0xce, 0x42, 0x5d, 0x7b,
	0x53, 0x97, 0xbe, 0x31, 0x48, 0xdf, 0xb5, 0xd7, 0xf5, 0x23, 0x14, 0xff, 0x3e, 0x0c, 0x4b, 0x2e,
	0x8a, 0x8b, 0xa3, 0x19, 0x76, 0x9e, 0xf7, 0xf2, 0x17, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x72, 0xaf,
	0x85, 0xef, 0x00, 0x00, 0x00,
}
