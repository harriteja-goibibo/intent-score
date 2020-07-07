// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/proto/common/common.proto

package common

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

type Org int32

const (
	Org_UNKNOWN_ORG Org = 0
	Org_GI          Org = 1
	Org_MMT         Org = 2
)

var Org_name = map[int32]string{
	0: "UNKNOWN_ORG",
	1: "GI",
	2: "MMT",
}

var Org_value = map[string]int32{
	"UNKNOWN_ORG": 0,
	"GI":          1,
	"MMT":         2,
}

func (x Org) String() string {
	return proto.EnumName(Org_name, int32(x))
}

func (Org) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5de68e6ee17ad41, []int{0}
}

type Vertical int32

const (
	Vertical_UNKNOWN_VERTICAL Vertical = 0
	Vertical_HOTELS           Vertical = 1
	Vertical_FLIGHTS          Vertical = 2
	Vertical_CARS             Vertical = 3
	Vertical_BUS              Vertical = 4
	Vertical_TRAINS           Vertical = 5
	Vertical_METRO            Vertical = 6
	Vertical_EXPERIENCES      Vertical = 7
	Vertical_HOLIDAYS         Vertical = 8
)

var Vertical_name = map[int32]string{
	0: "UNKNOWN_VERTICAL",
	1: "HOTELS",
	2: "FLIGHTS",
	3: "CARS",
	4: "BUS",
	5: "TRAINS",
	6: "METRO",
	7: "EXPERIENCES",
	8: "HOLIDAYS",
}

var Vertical_value = map[string]int32{
	"UNKNOWN_VERTICAL": 0,
	"HOTELS":           1,
	"FLIGHTS":          2,
	"CARS":             3,
	"BUS":              4,
	"TRAINS":           5,
	"METRO":            6,
	"EXPERIENCES":      7,
	"HOLIDAYS":         8,
}

func (x Vertical) String() string {
	return proto.EnumName(Vertical_name, int32(x))
}

func (Vertical) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5de68e6ee17ad41, []int{1}
}

type PageHitType int32

const (
	PageHitType_UNKNOWN_PAGE     PageHitType = 0
	PageHitType_SEARCH_PAGE      PageHitType = 1
	PageHitType_DETAIL_PAGE      PageHitType = 2
	PageHitType_REVIEW_PAGE      PageHitType = 3
	PageHitType_PAYMENT_PAGE     PageHitType = 4
	PageHitType_TRANSACTION_PAGE PageHitType = 5
)

var PageHitType_name = map[int32]string{
	0: "UNKNOWN_PAGE",
	1: "SEARCH_PAGE",
	2: "DETAIL_PAGE",
	3: "REVIEW_PAGE",
	4: "PAYMENT_PAGE",
	5: "TRANSACTION_PAGE",
}

var PageHitType_value = map[string]int32{
	"UNKNOWN_PAGE":     0,
	"SEARCH_PAGE":      1,
	"DETAIL_PAGE":      2,
	"REVIEW_PAGE":      3,
	"PAYMENT_PAGE":     4,
	"TRANSACTION_PAGE": 5,
}

func (x PageHitType) String() string {
	return proto.EnumName(PageHitType_name, int32(x))
}

func (PageHitType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f5de68e6ee17ad41, []int{2}
}

func init() {
	proto.RegisterEnum("common.Org", Org_name, Org_value)
	proto.RegisterEnum("common.Vertical", Vertical_name, Vertical_value)
	proto.RegisterEnum("common.PageHitType", PageHitType_name, PageHitType_value)
}

func init() { proto.RegisterFile("api/proto/common/common.proto", fileDescriptor_f5de68e6ee17ad41) }

var fileDescriptor_f5de68e6ee17ad41 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x91, 0x4d, 0x6e, 0xdb, 0x30,
	0x10, 0x85, 0xf5, 0x63, 0xcb, 0x2a, 0x6d, 0xa0, 0x03, 0xa2, 0xdb, 0x6a, 0x5b, 0xc0, 0x40, 0xad,
	0x45, 0x2f, 0x50, 0x5a, 0x66, 0x25, 0xa2, 0xfa, 0x03, 0x49, 0xdb, 0x75, 0x37, 0x85, 0x25, 0x08,
	0xaa, 0x90, 0xd8, 0x14, 0x14, 0x65, 0x91, 0x2c, 0x73, 0x8a, 0x1c, 0x29, 0xcb, 0x1c, 0x21, 0x70,
	0x2e, 0x12, 0xd0, 0xb2, 0x57, 0x04, 0x3f, 0xbc, 0x79, 0x78, 0xf3, 0x06, 0x7d, 0xdd, 0xb7, 0x8d,
	0xdf, 0x76, 0xaa, 0x57, 0x7e, 0xa9, 0x0e, 0x07, 0x75, 0xbc, 0x3c, 0x8b, 0x33, 0xc3, 0xce, 0xf0,
	0x9b, 0x7f, 0x43, 0x76, 0xd6, 0xd5, 0xf8, 0x33, 0x9a, 0xae, 0xd3, 0xdf, 0x69, 0xb6, 0x4d, 0xff,
	0x65, 0x3c, 0x04, 0x03, 0x3b, 0xc8, 0x0a, 0x19, 0x98, 0x78, 0x82, 0xec, 0x24, 0x91, 0x60, 0xcd,
	0x9f, 0x4c, 0xe4, 0x6e, 0xaa, 0xae, 0x6f, 0xca, 0xfd, 0x2d, 0xfe, 0x82, 0xe0, 0x2a, 0xdf, 0x50,
	0x2e, 0x59, 0x40, 0x62, 0x30, 0x30, 0x42, 0x4e, 0x94, 0x49, 0x1a, 0x0b, 0x30, 0xf1, 0x14, 0x4d,
	0x7e, 0xc5, 0x2c, 0x8c, 0xa4, 0x00, 0x0b, 0xbb, 0x68, 0x14, 0x10, 0x2e, 0xc0, 0xd6, 0x76, 0xcb,
	0xb5, 0x80, 0x91, 0xd6, 0x4a, 0x4e, 0x58, 0x2a, 0x60, 0x8c, 0x3f, 0xa1, 0x71, 0x42, 0x25, 0xcf,
	0xc0, 0xd1, 0x39, 0xe8, 0x9f, 0x9c, 0x72, 0x46, 0xd3, 0x80, 0x0a, 0x98, 0xe0, 0x19, 0x72, 0xa3,
	0x2c, 0x66, 0x2b, 0xb2, 0x13, 0xe0, 0xce, 0x1f, 0xd1, 0x34, 0xdf, 0xd7, 0x55, 0xd4, 0xf4, 0xf2,
	0xa1, 0xad, 0x30, 0xa0, 0xd9, 0x35, 0x46, 0x4e, 0x42, 0x0a, 0x86, 0x9e, 0x17, 0x94, 0xf0, 0x20,
	0x1a, 0x80, 0xa9, 0xc1, 0x8a, 0x4a, 0xc2, 0xe2, 0x01, 0x58, 0x1a, 0x70, 0xba, 0x61, 0x74, 0x3b,
	0x00, 0x5b, 0x9b, 0xe4, 0x64, 0x97, 0xd0, 0x54, 0x0e, 0x64, 0xa4, 0xb7, 0x93, 0x9c, 0xa4, 0x82,
	0x04, 0x92, 0x65, 0x17, 0xeb, 0xf1, 0xf2, 0xe7, 0xcb, 0xc9, 0x33, 0x5f, 0x4f, 0x9e, 0xf9, 0x76,
	0xf2, 0xcc, 0xe7, 0x77, 0xcf, 0xf8, 0xbb, 0xa8, 0x9b, 0xfe, 0xff, 0x7d, 0xb1, 0x28, 0xd5, 0xc1,
	0xaf, 0x55, 0x53, 0x34, 0x85, 0xf2, 0x9b, 0x63, 0x5f, 0x1d, 0xfb, 0xef, 0x77, 0xa5, 0xea, 0x2a,
	0xbf, 0xbd, 0xa9, 0x7d, 0x7d, 0x86, 0xa1, 0xeb, 0xc2, 0x39, 0x57, 0xff, 0xe3, 0x23, 0x00, 0x00,
	0xff, 0xff, 0x44, 0xaf, 0x17, 0x71, 0x9b, 0x01, 0x00, 0x00,
}