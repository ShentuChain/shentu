// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/oracle/v1alpha1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type GenesisState struct {
	Operators       []Operator                               `protobuf:"bytes,1,rep,name=operators,proto3" json:"operators" yaml:"operators"`
	TotalCollateral github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=total_collateral,json=totalCollateral,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_collateral" yaml:"total_collateral"`
	PoolParams      *LockedPoolParams                        `protobuf:"bytes,3,opt,name=pool_params,json=poolParams,proto3" json:"pool_params,omitempty" yaml:"pool_params"`
	TaskParams      *TaskParams                              `protobuf:"bytes,4,opt,name=task_params,json=taskParams,proto3" json:"task_params,omitempty" yaml:"task_params"`
	Withdraws       []Withdraw                               `protobuf:"bytes,5,rep,name=withdraws,proto3" json:"withdraws" yaml:"withdraws"`
	Tasks           []Task                                   `protobuf:"bytes,6,rep,name=tasks,proto3" json:"tasks" yaml:"tasks"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_6713fe00b3140e8c, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "shentu.oracle.v1alpha1.GenesisState")
}

func init() {
	proto.RegisterFile("shentu/oracle/v1alpha1/genesis.proto", fileDescriptor_6713fe00b3140e8c)
}

var fileDescriptor_6713fe00b3140e8c = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0x6d, 0xfa, 0x47, 0xe0, 0x54, 0xa2, 0xb2, 0xaa, 0x62, 0x2a, 0x64, 0x47, 0x07, 0x43,
	0x16, 0xee, 0x94, 0xb2, 0x75, 0x74, 0x07, 0x90, 0x8a, 0x44, 0x65, 0x90, 0x40, 0x30, 0x54, 0x6f,
	0x9c, 0x23, 0xb1, 0x7c, 0xf1, 0x6b, 0xf9, 0x2e, 0x94, 0x7e, 0x03, 0x46, 0xf8, 0x06, 0x9d, 0xf9,
	0x24, 0x1d, 0x3b, 0x32, 0x05, 0x94, 0x2c, 0x88, 0xb1, 0x9f, 0x00, 0xf9, 0xee, 0xe2, 0x44, 0x85,
	0x64, 0xb2, 0x25, 0xff, 0xee, 0xf7, 0xbc, 0xaf, 0x9f, 0xf3, 0x9e, 0xc8, 0x21, 0x2f, 0xd4, 0x98,
	0x61, 0x05, 0xa9, 0xe0, 0xec, 0x53, 0x17, 0x44, 0x39, 0x84, 0x2e, 0x1b, 0xf0, 0x82, 0xcb, 0x4c,
	0xd2, 0xb2, 0x42, 0x85, 0xfe, 0xbe, 0xa1, 0xa8, 0xa1, 0xe8, 0x9c, 0x3a, 0xd8, 0x1b, 0xe0, 0x00,
	0x35, 0xc2, 0xea, 0x37, 0x43, 0x1f, 0x84, 0x29, 0xca, 0x11, 0x4a, 0xd6, 0x03, 0x59, 0x1b, 0x7b,
	0x5c, 0x41, 0x97, 0xa5, 0x98, 0x15, 0xf6, 0xfb, 0xe3, 0x15, 0x99, 0xd6, 0xae, 0x21, 0xf2, 0x67,
	0xd3, 0xdb, 0x79, 0x6e, 0x86, 0x78, 0xad, 0x40, 0x71, 0xff, 0x9d, 0x77, 0x0f, 0x4b, 0x5e, 0x81,
	0xc2, 0x4a, 0x06, 0x6e, 0x7b, 0xa3, 0xd3, 0x3a, 0x6c, 0xd3, 0xff, 0xcf, 0x45, 0x5f, 0x59, 0x30,
	0x0e, 0xae, 0x26, 0x91, 0x73, 0x33, 0x89, 0x76, 0x2f, 0x60, 0x24, 0x8e, 0x48, 0x23, 0x20, 0xc9,
	0x42, 0xe6, 0x7f, 0x73, 0xbd, 0x5d, 0x85, 0x0a, 0xc4, 0x59, 0x8a, 0x42, 0x80, 0xe2, 0x15, 0x88,
	0xe0, 0x8e, 0x4e, 0x78, 0x48, 0xcd, 0x2e, 0xb4, 0xde, 0x85, 0xda, 0x5d, 0xe8, 0x31, 0x66, 0x45,
	0x7c, 0x62, 0xd5, 0x0f, 0x8c, 0xfa, 0xb6, 0x80, 0x7c, 0xff, 0x19, 0x75, 0x06, 0x99, 0x1a, 0x8e,
	0x7b, 0x34, 0xc5, 0x11, 0xb3, 0xff, 0xc4, 0x3c, 0x9e, 0xca, 0x7e, 0xce, 0xd4, 0x45, 0xc9, 0xa5,
	0x76, 0xc9, 0xe4, 0xbe, 0x3e, 0x7e, 0xdc, 0x9c, 0xf6, 0xc1, 0x6b, 0x95, 0x88, 0xe2, 0xac, 0x84,
	0x0a, 0x46, 0x32, 0xd8, 0x68, 0xbb, 0x9d, 0xd6, 0x61, 0x67, 0xd5, 0xbe, 0x2f, 0x31, 0xcd, 0x79,
	0xff, 0x14, 0x51, 0x9c, 0x6a, 0x3e, 0xde, 0xbf, 0x99, 0x44, 0xbe, 0x19, 0x6c, 0x49, 0x43, 0x12,
	0xaf, 0x6c, 0x18, 0xff, 0x83, 0xd7, 0x52, 0x20, 0xf3, 0x79, 0xc4, 0xa6, 0x8e, 0x20, 0xab, 0x22,
	0xde, 0x80, 0xcc, 0xff, 0x95, 0x2f, 0x09, 0x48, 0xe2, 0xa9, 0x86, 0xa9, 0xdb, 0x3a, 0xcf, 0xd4,
	0xb0, 0x5f, 0xc1, 0xb9, 0x0c, 0xb6, 0xd6, 0xb7, 0xf5, 0xd6, 0x82, 0xb7, 0xdb, 0x6a, 0x04, 0x24,
	0x59, 0xc8, 0xfc, 0x17, 0xde, 0x56, 0x9d, 0x23, 0x83, 0x6d, 0x6d, 0x7d, 0xb4, 0x6e, 0xe0, 0x78,
	0xcf, 0x1a, 0x77, 0x16, 0xe3, 0x4a, 0x92, 0x18, 0xc1, 0xd1, 0xdd, 0x2f, 0x97, 0x91, 0xf3, 0xfb,
	0x32, 0x72, 0xe2, 0x93, 0xab, 0x69, 0xe8, 0x5e, 0x4f, 0x43, 0xf7, 0xd7, 0x34, 0x74, 0xbf, 0xce,
	0x42, 0xe7, 0x7a, 0x16, 0x3a, 0x3f, 0x66, 0xa1, 0xf3, 0xbe, 0xbb, 0x5c, 0x21, 0xaf, 0x54, 0x96,
	0x7f, 0xc4, 0x71, 0xd1, 0x07, 0x95, 0x61, 0xc1, 0xec, 0x3d, 0xfe, 0x3c, 0xbf, 0xc9, 0xba, 0xd1,
	0xde, 0xb6, 0xbe, 0xc0, 0xcf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xf3, 0x12, 0xcf, 0x5b,
	0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tasks) > 0 {
		for iNdEx := len(m.Tasks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tasks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Withdraws) > 0 {
		for iNdEx := len(m.Withdraws) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Withdraws[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.TaskParams != nil {
		{
			size, err := m.TaskParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.PoolParams != nil {
		{
			size, err := m.PoolParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TotalCollateral) > 0 {
		for iNdEx := len(m.TotalCollateral) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TotalCollateral[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Operators) > 0 {
		for iNdEx := len(m.Operators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Operators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Operators) > 0 {
		for _, e := range m.Operators {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TotalCollateral) > 0 {
		for _, e := range m.TotalCollateral {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.PoolParams != nil {
		l = m.PoolParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TaskParams != nil {
		l = m.TaskParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Withdraws) > 0 {
		for _, e := range m.Withdraws {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Tasks) > 0 {
		for _, e := range m.Tasks {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operators", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Operators = append(m.Operators, Operator{})
			if err := m.Operators[len(m.Operators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCollateral", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TotalCollateral = append(m.TotalCollateral, types.Coin{})
			if err := m.TotalCollateral[len(m.TotalCollateral)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PoolParams == nil {
				m.PoolParams = &LockedPoolParams{}
			}
			if err := m.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TaskParams == nil {
				m.TaskParams = &TaskParams{}
			}
			if err := m.TaskParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdraws", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Withdraws = append(m.Withdraws, Withdraw{})
			if err := m.Withdraws[len(m.Withdraws)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tasks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tasks = append(m.Tasks, Task{})
			if err := m.Tasks[len(m.Tasks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
