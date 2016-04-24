// Code generated by protoc-gen-go.
// source: protot.proto
// DO NOT EDIT!

package protot

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type TextQueryType int32

const (
	TextQueryType_NOT_A_TEXT TextQueryType = 0
	TextQueryType_EXACT      TextQueryType = 1
	TextQueryType_HAS_PREFIX TextQueryType = 2
	TextQueryType_HAS_SUFFIX TextQueryType = 3
	TextQueryType_SUBSTRING  TextQueryType = 4
	TextQueryType_PATTERN    TextQueryType = 5
	TextQueryType_MIN_LENGTH TextQueryType = 6
	TextQueryType_MAX_LENGTH TextQueryType = 7
)

var TextQueryType_name = map[int32]string{
	0: "NOT_A_TEXT",
	1: "EXACT",
	2: "HAS_PREFIX",
	3: "HAS_SUFFIX",
	4: "SUBSTRING",
	5: "PATTERN",
	6: "MIN_LENGTH",
	7: "MAX_LENGTH",
}
var TextQueryType_value = map[string]int32{
	"NOT_A_TEXT": 0,
	"EXACT":      1,
	"HAS_PREFIX": 2,
	"HAS_SUFFIX": 3,
	"SUBSTRING":  4,
	"PATTERN":    5,
	"MIN_LENGTH": 6,
	"MAX_LENGTH": 7,
}

func (x TextQueryType) String() string {
	return proto.EnumName(TextQueryType_name, int32(x))
}
func (TextQueryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// NumericQueryType ...
type NumericQueryType int32

const (
	NumericQueryType_NOT_A_NUMBER  NumericQueryType = 0
	NumericQueryType_EQUAL         NumericQueryType = 1
	NumericQueryType_NOT_EQUAL     NumericQueryType = 2
	NumericQueryType_GREATER       NumericQueryType = 3
	NumericQueryType_GREATER_EQUAL NumericQueryType = 4
	NumericQueryType_LESS          NumericQueryType = 5
	NumericQueryType_LESS_EQUAL    NumericQueryType = 6
	NumericQueryType_IN            NumericQueryType = 7
	NumericQueryType_BETWEEN       NumericQueryType = 8
)

var NumericQueryType_name = map[int32]string{
	0: "NOT_A_NUMBER",
	1: "EQUAL",
	2: "NOT_EQUAL",
	3: "GREATER",
	4: "GREATER_EQUAL",
	5: "LESS",
	6: "LESS_EQUAL",
	7: "IN",
	8: "BETWEEN",
}
var NumericQueryType_value = map[string]int32{
	"NOT_A_NUMBER":  0,
	"EQUAL":         1,
	"NOT_EQUAL":     2,
	"GREATER":       3,
	"GREATER_EQUAL": 4,
	"LESS":          5,
	"LESS_EQUAL":    6,
	"IN":            7,
	"BETWEEN":       8,
}

func (x NumericQueryType) String() string {
	return proto.EnumName(NumericQueryType_name, int32(x))
}
func (NumericQueryType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// QueryString ...
type QueryString struct {
	Values   []string      `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	Valid    bool          `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Negation bool          `protobuf:"varint,3,opt,name=negation" json:"negation,omitempty"`
	Type     TextQueryType `protobuf:"varint,4,opt,name=type,enum=protot.TextQueryType" json:"type,omitempty"`
}

func (m *QueryString) Reset()                    { *m = QueryString{} }
func (m *QueryString) String() string            { return proto.CompactTextString(m) }
func (*QueryString) ProtoMessage()               {}
func (*QueryString) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// QueryInt64 ...
type QueryInt64 struct {
	Values   []int64          `protobuf:"varint,1,rep,name=values" json:"values,omitempty"`
	Valid    bool             `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Negation bool             `protobuf:"varint,3,opt,name=negation" json:"negation,omitempty"`
	Type     NumericQueryType `protobuf:"varint,4,opt,name=type,enum=protot.NumericQueryType" json:"type,omitempty"`
}

func (m *QueryInt64) Reset()                    { *m = QueryInt64{} }
func (m *QueryInt64) String() string            { return proto.CompactTextString(m) }
func (*QueryInt64) ProtoMessage()               {}
func (*QueryInt64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// QueryFloat64 ...
type QueryFloat64 struct {
	Values   []float64        `protobuf:"fixed64,1,rep,name=values" json:"values,omitempty"`
	Valid    bool             `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Negation bool             `protobuf:"varint,3,opt,name=negation" json:"negation,omitempty"`
	Type     NumericQueryType `protobuf:"varint,4,opt,name=type,enum=protot.NumericQueryType" json:"type,omitempty"`
}

func (m *QueryFloat64) Reset()                    { *m = QueryFloat64{} }
func (m *QueryFloat64) String() string            { return proto.CompactTextString(m) }
func (*QueryFloat64) ProtoMessage()               {}
func (*QueryFloat64) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// QueryTimestamp ...
type QueryTimestamp struct {
	Values   []*google_protobuf.Timestamp `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
	Valid    bool                         `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Negation bool                         `protobuf:"varint,3,opt,name=negation" json:"negation,omitempty"`
	Type     NumericQueryType             `protobuf:"varint,4,opt,name=type,enum=protot.NumericQueryType" json:"type,omitempty"`
}

func (m *QueryTimestamp) Reset()                    { *m = QueryTimestamp{} }
func (m *QueryTimestamp) String() string            { return proto.CompactTextString(m) }
func (*QueryTimestamp) ProtoMessage()               {}
func (*QueryTimestamp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryTimestamp) GetValues() []*google_protobuf.Timestamp {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryString)(nil), "protot.QueryString")
	proto.RegisterType((*QueryInt64)(nil), "protot.QueryInt64")
	proto.RegisterType((*QueryFloat64)(nil), "protot.QueryFloat64")
	proto.RegisterType((*QueryTimestamp)(nil), "protot.QueryTimestamp")
	proto.RegisterEnum("protot.TextQueryType", TextQueryType_name, TextQueryType_value)
	proto.RegisterEnum("protot.NumericQueryType", NumericQueryType_name, NumericQueryType_value)
}

var fileDescriptor0 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x93, 0x4d, 0x6b, 0xdb, 0x30,
	0x1c, 0xc6, 0xe7, 0xc4, 0x71, 0x9c, 0x7f, 0x5e, 0xd0, 0x04, 0x03, 0xd3, 0x53, 0xe9, 0x60, 0x94,
	0x1c, 0x1c, 0xe8, 0xc6, 0x2e, 0x3b, 0x39, 0x43, 0x49, 0x0d, 0xa9, 0xd6, 0xda, 0x0a, 0x0b, 0xbb,
	0x04, 0x27, 0xd3, 0x34, 0x43, 0x6c, 0x19, 0x47, 0x1e, 0xeb, 0x07, 0xd8, 0x6e, 0xfb, 0xce, 0xb3,
	0xfc, 0x42, 0x48, 0x77, 0xd8, 0xa5, 0x27, 0xe9, 0xe7, 0x47, 0xe8, 0xf7, 0x48, 0xc2, 0x30, 0xca,
	0x72, 0xa9, 0xa4, 0x72, 0xab, 0x01, 0x5b, 0x35, 0x5d, 0x7c, 0x10, 0xb1, 0xfa, 0x5e, 0xec, 0xdc,
	0xbd, 0x4c, 0x66, 0x42, 0x1e, 0xa2, 0x54, 0xcc, 0xaa, 0x64, 0x57, 0x7c, 0x9b, 0x65, 0xea, 0x31,
	0xe3, 0xc7, 0x99, 0x8a, 0x13, 0x7e, 0x54, 0x51, 0x92, 0x9d, 0x66, 0xf5, 0x26, 0x57, 0x7b, 0x18,
	0x3e, 0x14, 0x3c, 0x7f, 0x0c, 0x55, 0x1e, 0xa7, 0x02, 0x4f, 0xc0, 0xfa, 0x11, 0x1d, 0x0a, 0x7e,
	0x74, 0x8c, 0xcb, 0xee, 0xf5, 0x00, 0x8f, 0xa1, 0x57, 0x72, 0xfc, 0xd5, 0xe9, 0x5c, 0x1a, 0xd7,
	0x36, 0x46, 0x60, 0xa7, 0x5c, 0x44, 0x2a, 0x96, 0xa9, 0xd3, 0xad, 0xbe, 0xbc, 0x06, 0x53, 0x2b,
	0x1c, 0xb3, 0xa4, 0xc9, 0xcd, 0x2b, 0xb7, 0x69, 0xc8, 0xf8, 0x4f, 0x55, 0xed, 0xcb, 0xca, 0xf0,
	0x8a, 0x03, 0x54, 0xe0, 0xa7, 0xea, 0xfd, 0xbb, 0x27, 0x8e, 0xee, 0xff, 0x1d, 0x6f, 0xce, 0x1c,
	0x4e, 0xeb, 0xa0, 0x45, 0xc2, 0xf3, 0x78, 0x7f, 0xd2, 0x08, 0x18, 0x55, 0xb0, 0x38, 0xc8, 0xe8,
	0x5f, 0x91, 0xf1, 0x7c, 0xa2, 0x5f, 0x06, 0x4c, 0x6a, 0x6a, 0x6f, 0x13, 0x4f, 0xcf, 0x5c, 0xc3,
	0x9b, 0x0b, 0x57, 0x48, 0x29, 0x0e, 0xdc, 0x6d, 0x9f, 0xc2, 0x3d, 0xad, 0x7d, 0xae, 0x1e, 0xd3,
	0xdf, 0x06, 0x8c, 0xcf, 0x6e, 0xba, 0x3c, 0x32, 0xd0, 0x4f, 0x6c, 0xeb, 0x6d, 0x19, 0xd9, 0x30,
	0xf4, 0x02, 0x0f, 0xa0, 0x47, 0x36, 0xde, 0x47, 0x86, 0x0c, 0x1d, 0xdd, 0x7a, 0xe1, 0xf6, 0x3e,
	0x20, 0x0b, 0x7f, 0x83, 0x3a, 0x2d, 0x87, 0xeb, 0x85, 0x66, 0xfd, 0x0c, 0x83, 0x70, 0x3d, 0x0f,
	0x59, 0xe0, 0xd3, 0x25, 0x32, 0xf1, 0x10, 0xfa, 0xf7, 0x1e, 0x63, 0x24, 0xa0, 0xa8, 0xa7, 0xd7,
	0xde, 0xf9, 0x74, 0xbb, 0x22, 0x74, 0xc9, 0x6e, 0x91, 0x55, 0xb1, 0xb7, 0x69, 0xb9, 0x3f, 0xfd,
	0x63, 0x00, 0x7a, 0xda, 0xae, 0x3c, 0xd7, 0xa8, 0xee, 0x42, 0xd7, 0x77, 0x73, 0x12, 0x34, 0x6d,
	0x1e, 0xd6, 0xde, 0x0a, 0xe9, 0xb7, 0x18, 0xe8, 0xb0, 0xc6, 0x8e, 0xb6, 0x2d, 0x03, 0xe2, 0x95,
	0xba, 0xb2, 0xc9, 0x4b, 0x18, 0x37, 0xd0, 0xe4, 0x26, 0xb6, 0xc1, 0x5c, 0x91, 0x30, 0xac, 0xab,
	0xe8, 0x59, 0x93, 0x58, 0xd8, 0x82, 0x8e, 0x4f, 0x51, 0x5f, 0xef, 0x30, 0x27, 0xec, 0x33, 0x21,
	0x14, 0xd9, 0x73, 0xfb, 0x4b, 0xf3, 0x73, 0xec, 0xea, 0xf1, 0xed, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x59, 0x5f, 0x26, 0x45, 0x3b, 0x03, 0x00, 0x00,
}
