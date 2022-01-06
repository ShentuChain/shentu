// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: codec.proto

package testdata

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GroupMetadata struct {
	Description string                                        `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Admin       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=admin,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"admin,omitempty"`
}

func (m *GroupMetadata) Reset()         { *m = GroupMetadata{} }
func (m *GroupMetadata) String() string { return proto.CompactTextString(m) }
func (*GroupMetadata) ProtoMessage()    {}
func (*GroupMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_9610d574777ab505, []int{0}
}
func (m *GroupMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupMetadata.Merge(m, src)
}
func (m *GroupMetadata) XXX_Size() int {
	return m.Size()
}
func (m *GroupMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_GroupMetadata proto.InternalMessageInfo

func (m *GroupMetadata) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GroupMetadata) GetAdmin() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Admin
	}
	return nil
}

type GroupMember struct {
	Group  github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=group,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"group,omitempty"`
	Member github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=member,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"member,omitempty"`
	Weight uint64                                        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
}

func (m *GroupMember) Reset()         { *m = GroupMember{} }
func (m *GroupMember) String() string { return proto.CompactTextString(m) }
func (*GroupMember) ProtoMessage()    {}
func (*GroupMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_9610d574777ab505, []int{1}
}
func (m *GroupMember) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GroupMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GroupMember.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GroupMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GroupMember.Merge(m, src)
}
func (m *GroupMember) XXX_Size() int {
	return m.Size()
}
func (m *GroupMember) XXX_DiscardUnknown() {
	xxx_messageInfo_GroupMember.DiscardUnknown(m)
}

var xxx_messageInfo_GroupMember proto.InternalMessageInfo

func (m *GroupMember) GetGroup() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *GroupMember) GetMember() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Member
	}
	return nil
}

func (m *GroupMember) GetWeight() uint64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterType((*GroupMetadata)(nil), "testdata.GroupMetadata")
	proto.RegisterType((*GroupMember)(nil), "testdata.GroupMember")
}

func init() { proto.RegisterFile("codec.proto", fileDescriptor_9610d574777ab505) }

var fileDescriptor_9610d574777ab505 = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0xff, 0x87, 0x08, 0x1c, 0x58, 0x22, 0x84, 0x22, 0x06, 0x13, 0x75, 0xca, 0xd2,
	0x58, 0x08, 0x5e, 0xa0, 0x5d, 0x2a, 0x06, 0x18, 0x32, 0xb2, 0x25, 0xb6, 0x49, 0xac, 0x92, 0xdc,
	0xc8, 0xbe, 0x11, 0x82, 0xa7, 0xe0, 0x75, 0x78, 0x03, 0xc6, 0x8e, 0x4c, 0x08, 0x25, 0x6f, 0xc1,
	0x84, 0x9c, 0x06, 0xa9, 0x73, 0x27, 0xfb, 0x1e, 0xfb, 0x7c, 0xc7, 0xf2, 0xa1, 0x81, 0x00, 0xa9,
	0x44, 0xda, 0x1a, 0x40, 0x08, 0x8f, 0x50, 0x59, 0x94, 0x39, 0xe6, 0x17, 0x67, 0x25, 0x94, 0x30,
	0x8a, 0xdc, 0xed, 0xb6, 0xe7, 0xb3, 0x57, 0x7a, 0xba, 0x32, 0xd0, 0xb5, 0x77, 0x0a, 0x73, 0x77,
	0x2d, 0x8c, 0x69, 0x20, 0x95, 0x15, 0x46, 0xb7, 0xa8, 0xa1, 0x89, 0x48, 0x4c, 0x92, 0xe3, 0x6c,
	0x57, 0x0a, 0x57, 0xf4, 0x30, 0x97, 0xb5, 0x6e, 0xa2, 0x7f, 0x31, 0x49, 0x4e, 0x96, 0x57, 0x3f,
	0x5f, 0x97, 0xf3, 0x52, 0x63, 0xd5, 0x15, 0xa9, 0x80, 0x9a, 0x0b, 0xb0, 0x35, 0xd8, 0x69, 0x99,
	0x5b, 0xb9, 0xe6, 0xf8, 0xd2, 0x2a, 0x9b, 0x2e, 0x84, 0x58, 0x48, 0x69, 0x94, 0xb5, 0xd9, 0xd6,
	0x3f, 0x7b, 0x27, 0x34, 0x98, 0xc2, 0xeb, 0x42, 0x19, 0x07, 0x2e, 0xdd, 0x38, 0x86, 0xee, 0x07,
	0x1e, 0xfd, 0xe1, 0x2d, 0xf5, 0xeb, 0x11, 0xb9, 0xff, 0x13, 0x27, 0x40, 0x78, 0x4e, 0xfd, 0x67,
	0xa5, 0xcb, 0x0a, 0xa3, 0xff, 0x31, 0x49, 0x0e, 0xb2, 0x69, 0x5a, 0xde, 0x7f, 0xf4, 0x8c, 0x6c,
	0x7a, 0x46, 0xbe, 0x7b, 0x46, 0xde, 0x06, 0xe6, 0x6d, 0x06, 0xe6, 0x7d, 0x0e, 0xcc, 0x7b, 0xb8,
	0xd9, 0x0d, 0x52, 0x06, 0xf5, 0xfa, 0x11, 0xba, 0x46, 0xe6, 0xee, 0xef, 0xb8, 0xad, 0x54, 0x83,
	0x1d, 0x77, 0xa5, 0x74, 0xa8, 0x9f, 0xf8, 0x5f, 0x3b, 0x85, 0x3f, 0xd6, 0x71, 0xfd, 0x1b, 0x00,
	0x00, 0xff, 0xff, 0x9a, 0x96, 0x55, 0xd5, 0xbd, 0x01, 0x00, 0x00,
}

func (m *GroupMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GroupMember) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GroupMember) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GroupMember) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Weight != 0 {
		i = encodeVarintCodec(dAtA, i, uint64(m.Weight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Member) > 0 {
		i -= len(m.Member)
		copy(dAtA[i:], m.Member)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Member)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Group) > 0 {
		i -= len(m.Group)
		copy(dAtA[i:], m.Group)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Group)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	offset -= sovCodec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GroupMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func (m *GroupMember) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Group)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Member)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if m.Weight != 0 {
		n += 1 + sovCodec(uint64(m.Weight))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GroupMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: GroupMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = append(m.Admin[:0], dAtA[iNdEx:postIndex]...)
			if m.Admin == nil {
				m.Admin = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GroupMember) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: GroupMember: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GroupMember: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Group = append(m.Group[:0], dAtA[iNdEx:postIndex]...)
			if m.Group == nil {
				m.Group = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Member", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Member = append(m.Member[:0], dAtA[iNdEx:postIndex]...)
			if m.Member == nil {
				m.Member = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			m.Weight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Weight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCodec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCodec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCodec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCodec = fmt.Errorf("proto: unexpected end of group")
)
