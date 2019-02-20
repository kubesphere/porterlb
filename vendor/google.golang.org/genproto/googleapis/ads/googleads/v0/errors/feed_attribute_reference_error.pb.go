// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/feed_attribute_reference_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

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

// Enum describing possible feed attribute reference errors.
type FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError int32

const (
	// Enum unspecified.
	FeedAttributeReferenceErrorEnum_UNSPECIFIED FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 0
	// The received error code is not known in this version.
	FeedAttributeReferenceErrorEnum_UNKNOWN FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 1
	// A feed referenced by ID has been deleted.
	FeedAttributeReferenceErrorEnum_CANNOT_REFERENCE_DELETED_FEED FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 2
	// There is no active feed with the given name.
	FeedAttributeReferenceErrorEnum_INVALID_FEED_NAME FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 3
	// There is no feed attribute in an active feed with the given name.
	FeedAttributeReferenceErrorEnum_INVALID_FEED_ATTRIBUTE_NAME FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError = 4
)

var FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "CANNOT_REFERENCE_DELETED_FEED",
	3: "INVALID_FEED_NAME",
	4: "INVALID_FEED_ATTRIBUTE_NAME",
}
var FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_value = map[string]int32{
	"UNSPECIFIED":                   0,
	"UNKNOWN":                       1,
	"CANNOT_REFERENCE_DELETED_FEED": 2,
	"INVALID_FEED_NAME":             3,
	"INVALID_FEED_ATTRIBUTE_NAME":   4,
}

func (x FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError) String() string {
	return proto.EnumName(FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name, int32(x))
}
func (FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_feed_attribute_reference_error_a22b1b7f38948ef0, []int{0, 0}
}

// Container for enum describing possible feed attribute reference errors.
type FeedAttributeReferenceErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedAttributeReferenceErrorEnum) Reset()         { *m = FeedAttributeReferenceErrorEnum{} }
func (m *FeedAttributeReferenceErrorEnum) String() string { return proto.CompactTextString(m) }
func (*FeedAttributeReferenceErrorEnum) ProtoMessage()    {}
func (*FeedAttributeReferenceErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_feed_attribute_reference_error_a22b1b7f38948ef0, []int{0}
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Unmarshal(m, b)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Marshal(b, m, deterministic)
}
func (dst *FeedAttributeReferenceErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedAttributeReferenceErrorEnum.Merge(dst, src)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_Size() int {
	return xxx_messageInfo_FeedAttributeReferenceErrorEnum.Size(m)
}
func (m *FeedAttributeReferenceErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedAttributeReferenceErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedAttributeReferenceErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*FeedAttributeReferenceErrorEnum)(nil), "google.ads.googleads.v0.errors.FeedAttributeReferenceErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError", FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_name, FeedAttributeReferenceErrorEnum_FeedAttributeReferenceError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/feed_attribute_reference_error.proto", fileDescriptor_feed_attribute_reference_error_a22b1b7f38948ef0)
}

var fileDescriptor_feed_attribute_reference_error_a22b1b7f38948ef0 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xd1, 0x4e, 0xf2, 0x30,
	0x14, 0xc7, 0xbf, 0xc1, 0x17, 0x4d, 0xca, 0x85, 0x73, 0x89, 0x57, 0x44, 0x50, 0x1e, 0xa0, 0x5b,
	0xe2, 0xa5, 0x57, 0x65, 0x3b, 0x23, 0x8b, 0x58, 0xc8, 0x1c, 0x98, 0x98, 0x25, 0xcd, 0x60, 0xa5,
	0x21, 0x81, 0x95, 0xb4, 0x83, 0x47, 0xf1, 0x01, 0xbc, 0xf4, 0x11, 0x7c, 0x04, 0x1f, 0xc0, 0xe7,
	0x31, 0x5b, 0x1d, 0x89, 0x17, 0x72, 0xd5, 0x7f, 0x4e, 0x7f, 0xfd, 0x9d, 0xf6, 0x14, 0xf9, 0x42,
	0x4a, 0xb1, 0xe1, 0x6e, 0x96, 0x6b, 0xd7, 0xc4, 0x2a, 0x1d, 0x3c, 0x97, 0x2b, 0x25, 0x95, 0x76,
	0x57, 0x9c, 0xe7, 0x2c, 0x2b, 0x4b, 0xb5, 0x5e, 0xec, 0x4b, 0xce, 0x14, 0x5f, 0x71, 0xc5, 0x8b,
	0x25, 0x67, 0xf5, 0x3e, 0xde, 0x29, 0x59, 0x4a, 0xa7, 0x67, 0x4e, 0xe2, 0x2c, 0xd7, 0xf8, 0x28,
	0xc1, 0x07, 0x0f, 0x1b, 0xc9, 0xe0, 0xc3, 0x42, 0xfd, 0x90, 0xf3, 0x9c, 0x34, 0x9e, 0xb8, 0xd1,
	0x40, 0x05, 0x40, 0xb1, 0xdf, 0x0e, 0x5e, 0x2d, 0xd4, 0x3d, 0xc1, 0x38, 0x17, 0xa8, 0x33, 0xa3,
	0x4f, 0x53, 0xf0, 0xa3, 0x30, 0x82, 0xc0, 0xfe, 0xe7, 0x74, 0xd0, 0xf9, 0x8c, 0x3e, 0xd0, 0xc9,
	0x33, 0xb5, 0x2d, 0xe7, 0x16, 0x5d, 0xfb, 0x84, 0xd2, 0x49, 0xc2, 0x62, 0x08, 0x21, 0x06, 0xea,
	0x03, 0x0b, 0x60, 0x0c, 0x09, 0x04, 0x2c, 0x04, 0x08, 0xec, 0x96, 0x73, 0x85, 0x2e, 0x23, 0x3a,
	0x27, 0xe3, 0xc8, 0x54, 0x18, 0x25, 0x8f, 0x60, 0xb7, 0x9d, 0x3e, 0xea, 0xfe, 0x2a, 0x93, 0x24,
	0x89, 0xa3, 0xe1, 0x2c, 0x01, 0x03, 0xfc, 0x1f, 0x7e, 0x59, 0x68, 0xb0, 0x94, 0x5b, 0x7c, 0xfa,
	0x8d, 0xc3, 0x9b, 0x13, 0x97, 0x9f, 0x56, 0x53, 0x9a, 0x5a, 0x2f, 0xc1, 0x8f, 0x43, 0xc8, 0x4d,
	0x56, 0x08, 0x2c, 0x95, 0x70, 0x05, 0x2f, 0xea, 0x19, 0x36, 0xc3, 0xdf, 0xad, 0xf5, 0x5f, 0x7f,
	0x71, 0x6f, 0x96, 0xb7, 0x56, 0x7b, 0x44, 0xc8, 0x7b, 0xab, 0x37, 0x32, 0x32, 0x92, 0x6b, 0x6c,
	0x62, 0x95, 0xe6, 0x1e, 0xae, 0x5b, 0xea, 0xcf, 0x06, 0x48, 0x49, 0xae, 0xd3, 0x23, 0x90, 0xce,
	0xbd, 0xd4, 0x00, 0x8b, 0xb3, 0xba, 0xf1, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x51,
	0xca, 0x0c, 0x03, 0x02, 0x00, 0x00,
}
