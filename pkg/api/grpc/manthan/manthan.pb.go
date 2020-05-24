// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/proto/grpc/manthan/manthan.proto

package manthan

import (
	fmt "fmt"
	common "github.com/goibibo/intent-score/pkg/api/common"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type ManthanRealTimeData struct {
	Vertical common.Vertical    `protobuf:"varint,1,opt,name=vertical,proto3,enum=common.Vertical" json:"vertical,omitempty"`
	Org      common.Org         `protobuf:"varint,2,opt,name=org,proto3,enum=common.Org" json:"org,omitempty"`
	PageType common.PageHitType `protobuf:"varint,3,opt,name=page_type,json=pageType,proto3,enum=common.PageHitType" json:"page_type,omitempty"`
	// Format YYYYMMDDHHMMSS
	RequestDate string `protobuf:"bytes,4,opt,name=request_date,json=requestDate,proto3" json:"request_date,omitempty"`
	UserId      string `protobuf:"bytes,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Hotels, Experiences --> cityId
	// Bus, Flgits, Cars, Trains --> srcId-destId
	EntityId string `protobuf:"bytes,6,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// Format YYYYMMDD
	TravelStartDate string `protobuf:"bytes,7,opt,name=travel_start_date,json=travelStartDate,proto3" json:"travel_start_date,omitempty"`
	// Format YYYYMMDD
	TravelEndDate        string   `protobuf:"bytes,8,opt,name=travel_end_date,json=travelEndDate,proto3" json:"travel_end_date,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ManthanRealTimeData) Reset()         { *m = ManthanRealTimeData{} }
func (m *ManthanRealTimeData) String() string { return proto.CompactTextString(m) }
func (*ManthanRealTimeData) ProtoMessage()    {}
func (*ManthanRealTimeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c993833c6900966d, []int{0}
}
func (m *ManthanRealTimeData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManthanRealTimeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManthanRealTimeData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ManthanRealTimeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManthanRealTimeData.Merge(m, src)
}
func (m *ManthanRealTimeData) XXX_Size() int {
	return m.Size()
}
func (m *ManthanRealTimeData) XXX_DiscardUnknown() {
	xxx_messageInfo_ManthanRealTimeData.DiscardUnknown(m)
}

var xxx_messageInfo_ManthanRealTimeData proto.InternalMessageInfo

func (m *ManthanRealTimeData) GetVertical() common.Vertical {
	if m != nil {
		return m.Vertical
	}
	return common.Vertical_UNKNOWN_VERTICAL
}

func (m *ManthanRealTimeData) GetOrg() common.Org {
	if m != nil {
		return m.Org
	}
	return common.Org_UNKNOWN_ORG
}

func (m *ManthanRealTimeData) GetPageType() common.PageHitType {
	if m != nil {
		return m.PageType
	}
	return common.PageHitType_UNKNOWN_PAGE
}

func (m *ManthanRealTimeData) GetRequestDate() string {
	if m != nil {
		return m.RequestDate
	}
	return ""
}

func (m *ManthanRealTimeData) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ManthanRealTimeData) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func (m *ManthanRealTimeData) GetTravelStartDate() string {
	if m != nil {
		return m.TravelStartDate
	}
	return ""
}

func (m *ManthanRealTimeData) GetTravelEndDate() string {
	if m != nil {
		return m.TravelEndDate
	}
	return ""
}

func init() {
	proto.RegisterType((*ManthanRealTimeData)(nil), "manthan.ManthanRealTimeData")
}

func init() {
	proto.RegisterFile("api/proto/grpc/manthan/manthan.proto", fileDescriptor_c993833c6900966d)
}

var fileDescriptor_c993833c6900966d = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x41, 0x4e, 0xe3, 0x30,
	0x14, 0x86, 0x27, 0xed, 0x4c, 0x9b, 0xba, 0x33, 0x03, 0xb8, 0x0b, 0x22, 0x50, 0xa3, 0x82, 0x10,
	0xaa, 0x10, 0x34, 0x08, 0x38, 0x01, 0x2a, 0x88, 0x2e, 0x10, 0x28, 0x54, 0x2c, 0xd8, 0x44, 0x4e,
	0xf2, 0xe4, 0x5a, 0x34, 0x76, 0x70, 0x5e, 0x2b, 0xf5, 0x26, 0x5c, 0x87, 0x1d, 0x4b, 0x8e, 0x80,
	0xca, 0x45, 0x90, 0x9d, 0x14, 0x58, 0xbd, 0xbc, 0xef, 0xff, 0x62, 0xcb, 0xfa, 0xc9, 0x1e, 0xcb,
	0x45, 0x90, 0x6b, 0x85, 0x2a, 0xe0, 0x3a, 0x4f, 0x82, 0x8c, 0x49, 0x9c, 0x30, 0xb9, 0x9a, 0x03,
	0x1b, 0xd1, 0x66, 0xb5, 0x6e, 0x75, 0xbf, 0xf5, 0x44, 0x65, 0x99, 0x92, 0xd5, 0x28, 0xbd, 0xdd,
	0x97, 0x1a, 0xe9, 0x5c, 0x97, 0x6a, 0x08, 0x6c, 0x3a, 0x16, 0x19, 0x0c, 0x19, 0x32, 0x7a, 0x48,
	0xdc, 0x39, 0x68, 0x14, 0x09, 0x9b, 0x7a, 0x4e, 0xcf, 0xe9, 0xff, 0x3f, 0x59, 0x1f, 0x54, 0x3f,
	0xde, 0x57, 0x3c, 0xfc, 0x32, 0x68, 0x97, 0xd4, 0x95, 0xe6, 0x5e, 0xcd, 0x8a, 0xed, 0x95, 0x78,
	0xa3, 0x79, 0x68, 0x38, 0x3d, 0x26, 0xad, 0x9c, 0x71, 0x88, 0x70, 0x91, 0x83, 0x57, 0xb7, 0x52,
	0x67, 0x25, 0xdd, 0x32, 0x0e, 0x57, 0x02, 0xc7, 0x8b, 0x1c, 0x42, 0xd7, 0x58, 0xe6, 0x8b, 0xee,
	0x90, 0xbf, 0x1a, 0x9e, 0x66, 0x50, 0x60, 0x94, 0x32, 0x04, 0xef, 0x77, 0xcf, 0xe9, 0xb7, 0xc2,
	0x76, 0xc5, 0x86, 0x0c, 0x81, 0x6e, 0x92, 0xe6, 0xac, 0x00, 0x1d, 0x89, 0xd4, 0xfb, 0x63, 0xd3,
	0x86, 0x59, 0x47, 0x29, 0xdd, 0x26, 0x2d, 0x90, 0x28, 0x70, 0x61, 0xa2, 0x86, 0x8d, 0xdc, 0x12,
	0x8c, 0x52, 0x7a, 0x40, 0x36, 0x50, 0xb3, 0x39, 0x4c, 0xa3, 0x02, 0x99, 0xae, 0x4e, 0x6f, 0x5a,
	0x69, 0xad, 0x0c, 0xee, 0x0c, 0xb7, 0x37, 0xec, 0x93, 0x0a, 0x45, 0x20, 0xd3, 0xd2, 0x74, 0xad,
	0xf9, 0xaf, 0xc4, 0x17, 0x32, 0x35, 0xde, 0xf9, 0xe5, 0xeb, 0xd2, 0x77, 0xde, 0x96, 0xbe, 0xf3,
	0xbe, 0xf4, 0x9d, 0xe7, 0x0f, 0xff, 0xd7, 0xc3, 0x19, 0x17, 0x38, 0x99, 0xc5, 0xe6, 0x8d, 0x01,
	0x57, 0x22, 0x16, 0xb1, 0x0a, 0x84, 0x44, 0x90, 0x78, 0x54, 0x24, 0x4a, 0x43, 0x90, 0x3f, 0xf2,
	0xc0, 0x14, 0xf3, 0xb3, 0xc1, 0xb8, 0x61, 0x2b, 0x39, 0xfd, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x7d,
	0x42, 0x4c, 0x0a, 0xe2, 0x01, 0x00, 0x00,
}

func (m *ManthanRealTimeData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManthanRealTimeData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ManthanRealTimeData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.TravelEndDate) > 0 {
		i -= len(m.TravelEndDate)
		copy(dAtA[i:], m.TravelEndDate)
		i = encodeVarintManthan(dAtA, i, uint64(len(m.TravelEndDate)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.TravelStartDate) > 0 {
		i -= len(m.TravelStartDate)
		copy(dAtA[i:], m.TravelStartDate)
		i = encodeVarintManthan(dAtA, i, uint64(len(m.TravelStartDate)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.EntityId) > 0 {
		i -= len(m.EntityId)
		copy(dAtA[i:], m.EntityId)
		i = encodeVarintManthan(dAtA, i, uint64(len(m.EntityId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.UserId) > 0 {
		i -= len(m.UserId)
		copy(dAtA[i:], m.UserId)
		i = encodeVarintManthan(dAtA, i, uint64(len(m.UserId)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.RequestDate) > 0 {
		i -= len(m.RequestDate)
		copy(dAtA[i:], m.RequestDate)
		i = encodeVarintManthan(dAtA, i, uint64(len(m.RequestDate)))
		i--
		dAtA[i] = 0x22
	}
	if m.PageType != 0 {
		i = encodeVarintManthan(dAtA, i, uint64(m.PageType))
		i--
		dAtA[i] = 0x18
	}
	if m.Org != 0 {
		i = encodeVarintManthan(dAtA, i, uint64(m.Org))
		i--
		dAtA[i] = 0x10
	}
	if m.Vertical != 0 {
		i = encodeVarintManthan(dAtA, i, uint64(m.Vertical))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintManthan(dAtA []byte, offset int, v uint64) int {
	offset -= sovManthan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ManthanRealTimeData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vertical != 0 {
		n += 1 + sovManthan(uint64(m.Vertical))
	}
	if m.Org != 0 {
		n += 1 + sovManthan(uint64(m.Org))
	}
	if m.PageType != 0 {
		n += 1 + sovManthan(uint64(m.PageType))
	}
	l = len(m.RequestDate)
	if l > 0 {
		n += 1 + l + sovManthan(uint64(l))
	}
	l = len(m.UserId)
	if l > 0 {
		n += 1 + l + sovManthan(uint64(l))
	}
	l = len(m.EntityId)
	if l > 0 {
		n += 1 + l + sovManthan(uint64(l))
	}
	l = len(m.TravelStartDate)
	if l > 0 {
		n += 1 + l + sovManthan(uint64(l))
	}
	l = len(m.TravelEndDate)
	if l > 0 {
		n += 1 + l + sovManthan(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovManthan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozManthan(x uint64) (n int) {
	return sovManthan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ManthanRealTimeData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowManthan
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ManthanRealTimeData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManthanRealTimeData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vertical", wireType)
			}
			m.Vertical = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManthan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vertical |= common.Vertical(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Org", wireType)
			}
			m.Org = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManthan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Org |= common.Org(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageType", wireType)
			}
			m.PageType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManthan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageType |= common.PageHitType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManthan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManthan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthManthan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManthan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManthan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthManthan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManthan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManthan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthManthan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EntityId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TravelStartDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManthan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManthan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthManthan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TravelStartDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TravelEndDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowManthan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthManthan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthManthan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TravelEndDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipManthan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthManthan
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthManthan
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipManthan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowManthan
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowManthan
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowManthan
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthManthan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupManthan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthManthan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthManthan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowManthan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupManthan = fmt.Errorf("proto: unexpected end of group")
)
