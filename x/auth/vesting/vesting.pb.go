// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/vesting/v1alpha1/vesting.proto

package vesting

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
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

// ManualVestingAccount implements the VestingAccount interface.
type ManualVestingAccount struct {
	*types.BaseVestingAccount `protobuf:"bytes,1,opt,name=base_vesting_account,json=baseVestingAccount,proto3,embedded=base_vesting_account" json:"base_vesting_account,omitempty"`
	VestedCoins               github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=vested_coins,json=vestedCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"vested_coins" yaml:"vested_coins"`
	Unlocker                  string                                   `protobuf:"bytes,3,opt,name=unlocker,proto3" json:"unlocker,omitempty" yaml:"unlocker"`
}

func (m *ManualVestingAccount) Reset()      { *m = ManualVestingAccount{} }
func (*ManualVestingAccount) ProtoMessage() {}
func (*ManualVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_3367d243850d7f5d, []int{0}
}
func (m *ManualVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ManualVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ManualVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ManualVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ManualVestingAccount.Merge(m, src)
}
func (m *ManualVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *ManualVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ManualVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ManualVestingAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ManualVestingAccount)(nil), "cosmos.vesting.v1beta1.ManualVestingAccount")
}

func init() {
	proto.RegisterFile("shentu/vesting/v1alpha1/vesting.proto", fileDescriptor_3367d243850d7f5d)
}

var fileDescriptor_3367d243850d7f5d = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xb1, 0x4e, 0xc2, 0x40,
	0x18, 0xc7, 0x7b, 0x60, 0x0c, 0x16, 0x13, 0x93, 0x42, 0x0c, 0x32, 0xb4, 0xa4, 0xd1, 0xa4, 0x31,
	0xb1, 0x97, 0xe2, 0xc6, 0x66, 0x1d, 0x1c, 0x8c, 0x0b, 0x83, 0x83, 0x0b, 0xb9, 0x96, 0xb3, 0x6d,
	0x28, 0x77, 0x84, 0xbb, 0x23, 0xf2, 0x00, 0x26, 0x8e, 0x8e, 0x8e, 0xcc, 0xbe, 0x83, 0x3b, 0x23,
	0xa3, 0x53, 0x35, 0xf0, 0x06, 0x3c, 0x81, 0x69, 0xaf, 0x45, 0x88, 0x4e, 0xed, 0xfd, 0xff, 0xbf,
	0xef, 0xff, 0x7d, 0x97, 0xef, 0xd4, 0x33, 0x16, 0x62, 0xc2, 0x05, 0x9c, 0x60, 0xc6, 0x23, 0x12,
	0xc0, 0x89, 0x83, 0xe2, 0x51, 0x88, 0x9c, 0x42, 0xb0, 0x47, 0x63, 0xca, 0xa9, 0x76, 0xec, 0x53,
	0x36, 0xa4, 0xcc, 0x2e, 0xd4, 0x89, 0xe3, 0x61, 0x8e, 0x9c, 0x66, 0x3d, 0xa0, 0x01, 0xcd, 0x10,
	0x98, 0xfe, 0x49, 0xba, 0xa9, 0x4b, 0x1a, 0x7a, 0x88, 0x61, 0x98, 0xa3, 0xd0, 0xa7, 0x11, 0xc9,
	0xfd, 0xd3, 0xdc, 0xff, 0x6d, 0x2a, 0x91, 0x9d, 0x9e, 0xe6, 0x47, 0x49, 0xad, 0xdf, 0x21, 0x22,
	0x50, 0x7c, 0x2f, 0xf5, 0x2b, 0xdf, 0xa7, 0x82, 0x70, 0xcd, 0x53, 0xeb, 0x69, 0x72, 0x2f, 0xc7,
	0x7b, 0x48, 0xea, 0x0d, 0xd0, 0x02, 0x56, 0xb5, 0x7d, 0x6e, 0xff, 0x3f, 0xab, 0xed, 0x22, 0x86,
	0x77, 0x93, 0xdc, 0xbd, 0x45, 0x62, 0x80, 0xae, 0xe6, 0xfd, 0x71, 0xb4, 0x67, 0xa0, 0x1e, 0xa6,
	0x01, 0xb8, 0xdf, 0x4b, 0x07, 0x67, 0x8d, 0x52, 0xab, 0x6c, 0x55, 0xdb, 0x27, 0x45, 0x78, 0x5a,
	0xb2, 0x49, 0xbe, 0xa6, 0x11, 0x71, 0x6f, 0xe6, 0x89, 0xa1, 0xac, 0x13, 0xa3, 0x36, 0x45, 0xc3,
	0xb8, 0x63, 0x6e, 0x17, 0x9b, 0xef, 0x5f, 0x86, 0x15, 0x44, 0x3c, 0x14, 0x9e, 0xed, 0xd3, 0x21,
	0xcc, 0xaf, 0x2f, 0x3f, 0x17, 0xac, 0x3f, 0x80, 0x7c, 0x3a, 0xc2, 0x2c, 0xcb, 0x61, 0xdd, 0xaa,
	0x2c, 0xcd, 0x0e, 0x1a, 0x54, 0x2b, 0x82, 0xc4, 0xd4, 0x1f, 0xe0, 0x71, 0xa3, 0xdc, 0x02, 0xd6,
	0x81, 0x5b, 0x5b, 0x27, 0xc6, 0x91, 0xec, 0x51, 0x38, 0x66, 0x77, 0x03, 0x75, 0x2a, 0x2f, 0x33,
	0x43, 0x79, 0x9b, 0x19, 0x8a, 0x7b, 0x3b, 0x5f, 0xea, 0x60, 0xb1, 0xd4, 0xc1, 0xf7, 0x52, 0x07,
	0xaf, 0x2b, 0x5d, 0x59, 0xac, 0x74, 0xe5, 0x73, 0xa5, 0x2b, 0x0f, 0xce, 0xf6, 0x2c, 0x78, 0xcc,
	0xa3, 0xc1, 0x23, 0x15, 0xa4, 0x8f, 0x78, 0x44, 0x09, 0xcc, 0x1f, 0xc4, 0x13, 0x44, 0x82, 0x87,
	0xc5, 0x4a, 0xbc, 0xfd, 0x6c, 0x27, 0x97, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x82, 0xda, 0x98,
	0x4b, 0x30, 0x02, 0x00, 0x00,
}

func (m *ManualVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ManualVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ManualVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Unlocker) > 0 {
		i -= len(m.Unlocker)
		copy(dAtA[i:], m.Unlocker)
		i = encodeVarintVesting(dAtA, i, uint64(len(m.Unlocker)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VestedCoins) > 0 {
		for iNdEx := len(m.VestedCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestedCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintVesting(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BaseVestingAccount != nil {
		{
			size, err := m.BaseVestingAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVesting(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVesting(dAtA []byte, offset int, v uint64) int {
	offset -= sovVesting(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ManualVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseVestingAccount != nil {
		l = m.BaseVestingAccount.Size()
		n += 1 + l + sovVesting(uint64(l))
	}
	if len(m.VestedCoins) > 0 {
		for _, e := range m.VestedCoins {
			l = e.Size()
			n += 1 + l + sovVesting(uint64(l))
		}
	}
	l = len(m.Unlocker)
	if l > 0 {
		n += 1 + l + sovVesting(uint64(l))
	}
	return n
}

func sovVesting(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVesting(x uint64) (n int) {
	return sovVesting(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ManualVestingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVesting
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
			return fmt.Errorf("proto: ManualVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ManualVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseVestingAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseVestingAccount == nil {
				m.BaseVestingAccount = &types.BaseVestingAccount{}
			}
			if err := m.BaseVestingAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestedCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestedCoins = append(m.VestedCoins, types1.Coin{})
			if err := m.VestedCoins[len(m.VestedCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unlocker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVesting
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
				return ErrInvalidLengthVesting
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVesting
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unlocker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVesting(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVesting
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVesting
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
func skipVesting(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
					return 0, ErrIntOverflowVesting
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
				return 0, ErrInvalidLengthVesting
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVesting
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVesting
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVesting        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVesting          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVesting = fmt.Errorf("proto: unexpected end of group")
)
