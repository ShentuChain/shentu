// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/auth/v1alpha1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgUnlock defines a message for unlocking coins from a manual vesting
// account.
type MsgUnlock struct {
	Issuer       string                                   `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty" yaml:"issuer"`
	Account      string                                   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty" yaml:"account_address"`
	UnlockAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=unlock_amount,json=unlockAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"unlock_amount"`
}

func (m *MsgUnlock) Reset()         { *m = MsgUnlock{} }
func (m *MsgUnlock) String() string { return proto.CompactTextString(m) }
func (*MsgUnlock) ProtoMessage()    {}
func (*MsgUnlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e193e7ad0fae544, []int{0}
}
func (m *MsgUnlock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnlock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnlock.Merge(m, src)
}
func (m *MsgUnlock) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnlock) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnlock.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnlock proto.InternalMessageInfo

// MsgUnlockResponse defines the Msg/Unlock response type.
type MsgUnlockResponse struct {
}

func (m *MsgUnlockResponse) Reset()         { *m = MsgUnlockResponse{} }
func (m *MsgUnlockResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUnlockResponse) ProtoMessage()    {}
func (*MsgUnlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e193e7ad0fae544, []int{1}
}
func (m *MsgUnlockResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnlockResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnlockResponse.Merge(m, src)
}
func (m *MsgUnlockResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnlockResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUnlock)(nil), "shentu.auth.v1alpha1.MsgUnlock")
	proto.RegisterType((*MsgUnlockResponse)(nil), "shentu.auth.v1alpha1.MsgUnlockResponse")
}

func init() { proto.RegisterFile("shentu/auth/v1alpha1/tx.proto", fileDescriptor_2e193e7ad0fae544) }

var fileDescriptor_2e193e7ad0fae544 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x6e, 0xea, 0x30,
	0x14, 0xc6, 0x93, 0x8b, 0xc4, 0xbd, 0xf8, 0x96, 0x81, 0x14, 0x55, 0x34, 0x52, 0x13, 0x94, 0xa5,
	0x74, 0xa8, 0xdd, 0xd0, 0x4e, 0x6c, 0xa5, 0x53, 0x07, 0x96, 0x48, 0x1d, 0xda, 0x05, 0x39, 0x89,
	0x9b, 0x44, 0x10, 0x3b, 0xc2, 0x0e, 0x82, 0x37, 0xe8, 0xd8, 0x47, 0x60, 0xee, 0x93, 0x30, 0x32,
	0x76, 0xa2, 0x15, 0x2c, 0xcc, 0x3c, 0x41, 0x85, 0x1d, 0x50, 0x87, 0x4a, 0x9d, 0x6c, 0x9d, 0xef,
	0x77, 0xfe, 0x7c, 0xe7, 0x80, 0x33, 0x1e, 0x13, 0x2a, 0x72, 0x84, 0x73, 0x11, 0xa3, 0xb1, 0x8b,
	0x87, 0x59, 0x8c, 0x5d, 0x24, 0x26, 0x30, 0x1b, 0x31, 0xc1, 0x8c, 0xba, 0x92, 0xe1, 0x4e, 0x86,
	0x7b, 0xd9, 0xac, 0x47, 0x2c, 0x62, 0x12, 0x40, 0xbb, 0x9f, 0x62, 0x4d, 0x2b, 0x60, 0x3c, 0x65,
	0x1c, 0xf9, 0x98, 0x13, 0x34, 0x76, 0x7d, 0x22, 0xb0, 0x8b, 0x02, 0x96, 0x50, 0xa5, 0x3b, 0x1b,
	0x1d, 0x54, 0x7a, 0x3c, 0x7a, 0xa0, 0x43, 0x16, 0x0c, 0x8c, 0x0b, 0x50, 0x4e, 0x38, 0xcf, 0xc9,
	0xa8, 0xa1, 0x37, 0xf5, 0x56, 0xa5, 0x5b, 0xdb, 0x2e, 0xed, 0xea, 0x14, 0xa7, 0xc3, 0x8e, 0xa3,
	0xe2, 0x8e, 0x57, 0x00, 0xc6, 0x0d, 0xf8, 0x8b, 0x83, 0x80, 0xe5, 0x54, 0x34, 0xfe, 0x48, 0xd6,
	0xdc, 0x2e, 0xed, 0x13, 0xc5, 0x16, 0x42, 0x1f, 0x87, 0xe1, 0x88, 0x70, 0xee, 0x78, 0x7b, 0xd4,
	0xc8, 0x40, 0x35, 0x97, 0xad, 0xfa, 0x38, 0x95, 0xb9, 0xa5, 0x66, 0xa9, 0xf5, 0xbf, 0x7d, 0x0a,
	0xd5, 0x98, 0x70, 0x37, 0x26, 0x2c, 0xc6, 0x84, 0x77, 0x2c, 0xa1, 0xdd, 0xab, 0xf9, 0xd2, 0xd6,
	0xde, 0x3e, 0xec, 0x56, 0x94, 0x88, 0x38, 0xf7, 0x61, 0xc0, 0x52, 0x54, 0x78, 0x52, 0xcf, 0x25,
	0x0f, 0x07, 0x48, 0x4c, 0x33, 0xc2, 0x65, 0x02, 0xf7, 0x8e, 0x54, 0x87, 0x5b, 0xd9, 0xa0, 0xf3,
	0xef, 0x65, 0x66, 0x6b, 0x9b, 0x99, 0xad, 0x39, 0xc7, 0xa0, 0x76, 0x70, 0xea, 0x11, 0x9e, 0x31,
	0xca, 0x49, 0xfb, 0x11, 0x94, 0x7a, 0x3c, 0x32, 0x3c, 0x50, 0x2e, 0x56, 0x60, 0xc3, 0x9f, 0xb6,
	0x0b, 0x0f, 0x99, 0xe6, 0xf9, 0x2f, 0xc0, 0xbe, 0x74, 0xf7, 0x7e, 0xbe, 0xb2, 0xf4, 0xc5, 0xca,
	0xd2, 0x3f, 0x57, 0x96, 0xfe, 0xba, 0xb6, 0xb4, 0xc5, 0xda, 0xd2, 0xde, 0xd7, 0x96, 0xf6, 0x84,
	0xbe, 0x79, 0x51, 0xc5, 0x9e, 0x59, 0x4e, 0x43, 0x2c, 0x12, 0x46, 0x8b, 0x00, 0x9a, 0xa8, 0xeb,
	0x4b, 0x63, 0x7e, 0x59, 0x1e, 0xeb, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x64, 0x3f, 0xa2, 0x81,
	0x19, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Unlock defines a method for unlocking coins from a manual vesting
	// account.
	Unlock(ctx context.Context, in *MsgUnlock, opts ...grpc.CallOption) (*MsgUnlockResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Unlock(ctx context.Context, in *MsgUnlock, opts ...grpc.CallOption) (*MsgUnlockResponse, error) {
	out := new(MsgUnlockResponse)
	err := c.cc.Invoke(ctx, "/shentu.auth.v1alpha1.Msg/Unlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Unlock defines a method for unlocking coins from a manual vesting
	// account.
	Unlock(context.Context, *MsgUnlock) (*MsgUnlockResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Unlock(ctx context.Context, req *MsgUnlock) (*MsgUnlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unlock not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Unlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnlock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Unlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.auth.v1alpha1.Msg/Unlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Unlock(ctx, req.(*MsgUnlock))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shentu.auth.v1alpha1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Unlock",
			Handler:    _Msg_Unlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shentu/auth/v1alpha1/tx.proto",
}

func (m *MsgUnlock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnlock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnlock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UnlockAmount) > 0 {
		for iNdEx := len(m.UnlockAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UnlockAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnlockResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnlockResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnlockResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUnlock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.UnlockAmount) > 0 {
		for _, e := range m.UnlockAmount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgUnlockResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUnlock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUnlock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnlock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnlockAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UnlockAmount = append(m.UnlockAmount, types.Coin{})
			if err := m.UnlockAmount[len(m.UnlockAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUnlockResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUnlockResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnlockResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
