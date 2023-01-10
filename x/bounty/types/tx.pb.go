// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/bounty/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"

	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgCreateProgram defines a SDK message for creating a new validator.
type MsgCreateProgram struct {
	Description       string                                 `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	CommissionRate    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=commission_rate,json=commissionRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"commission_rate"`
	SubmissionEndTime time.Time                              `protobuf:"bytes,3,opt,name=submission_end_time,json=submissionEndTime,proto3,stdtime" json:"submission_end_time" yaml:"submission_end_time"`
	CreatorAddress    string                                 `protobuf:"bytes,4,opt,name=creator_address,json=creatorAddress,proto3" json:"creator_address,omitempty" yaml:"creator_address"`
	EncryptionKey     *types.Any                             `protobuf:"bytes,5,opt,name=encryption_key,json=encryptionKey,proto3" json:"encryption_key,omitempty" yaml:"encryption_key"`
	Deposit           []types1.Coin                          `protobuf:"bytes,6,rep,name=deposit,proto3" json:"deposit"`
}

func (m *MsgCreateProgram) Reset()         { *m = MsgCreateProgram{} }
func (m *MsgCreateProgram) String() string { return proto.CompactTextString(m) }
func (*MsgCreateProgram) ProtoMessage()    {}
func (*MsgCreateProgram) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4b4296bac3db30, []int{0}
}
func (m *MsgCreateProgram) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateProgram) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateProgram.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateProgram) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateProgram.Merge(m, src)
}
func (m *MsgCreateProgram) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateProgram) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateProgram.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateProgram proto.InternalMessageInfo

// MsgCreateProgramResponse defines the Msg/CreateValidator response type.
type MsgCreateProgramResponse struct {
	ProgramId uint64 `protobuf:"varint,1,opt,name=program_id,json=programId,proto3" json:"id" yaml:"id"`
}

func (m *MsgCreateProgramResponse) Reset()         { *m = MsgCreateProgramResponse{} }
func (m *MsgCreateProgramResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateProgramResponse) ProtoMessage()    {}
func (*MsgCreateProgramResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4b4296bac3db30, []int{1}
}
func (m *MsgCreateProgramResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateProgramResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateProgramResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateProgramResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateProgramResponse.Merge(m, src)
}
func (m *MsgCreateProgramResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateProgramResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateProgramResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateProgramResponse proto.InternalMessageInfo

func (m *MsgCreateProgramResponse) GetProgramId() uint64 {
	if m != nil {
		return m.ProgramId
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreateProgram)(nil), "shentu.bounty.v1.MsgCreateProgram")
	proto.RegisterType((*MsgCreateProgramResponse)(nil), "shentu.bounty.v1.MsgCreateProgramResponse")
}

func init() { proto.RegisterFile("shentu/bounty/v1/tx.proto", fileDescriptor_1e4b4296bac3db30) }

var fileDescriptor_1e4b4296bac3db30 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xcf, 0x6a, 0xd4, 0x40,
	0x1c, 0x4e, 0xba, 0xb5, 0xda, 0x29, 0xfd, 0x63, 0x5a, 0x25, 0x0d, 0x98, 0x2c, 0x39, 0x94, 0xa5,
	0xd0, 0x09, 0xbb, 0x9e, 0xec, 0xad, 0x59, 0x7b, 0x90, 0x52, 0x91, 0x20, 0x08, 0x5e, 0x42, 0x92,
	0x99, 0x4d, 0x87, 0x36, 0x33, 0x21, 0x33, 0x59, 0x9a, 0x37, 0xf0, 0xd8, 0x47, 0xe8, 0x43, 0xf8,
	0x10, 0xc5, 0x53, 0x8f, 0xe2, 0x21, 0xca, 0xee, 0x45, 0x3c, 0x2e, 0x3e, 0x80, 0x24, 0x93, 0xd8,
	0xed, 0x56, 0xf0, 0x94, 0xcc, 0xef, 0xfb, 0x7e, 0xbf, 0xf9, 0xe6, 0x9b, 0x6f, 0xc0, 0x2e, 0x3f,
	0xc3, 0x54, 0xe4, 0x4e, 0xc8, 0x72, 0x2a, 0x0a, 0x67, 0xdc, 0x77, 0xc4, 0x25, 0x4c, 0x33, 0x26,
	0x98, 0xb6, 0x25, 0x21, 0x28, 0x21, 0x38, 0xee, 0x1b, 0x3b, 0x31, 0x8b, 0x59, 0x0d, 0x3a, 0xd5,
	0x9f, 0xe4, 0x19, 0x56, 0xcc, 0x58, 0x7c, 0x81, 0x9d, 0x7a, 0x15, 0xe6, 0x23, 0x47, 0x90, 0x04,
	0x73, 0x11, 0x24, 0x69, 0x43, 0xd8, 0x8d, 0x18, 0x4f, 0x18, 0xf7, 0x65, 0xa7, 0x5c, 0xb4, 0xd0,
	0x62, 0x6f, 0x40, 0x8b, 0x06, 0x32, 0x25, 0xd1, 0x09, 0x03, 0x8e, 0x9d, 0x71, 0x3f, 0xc4, 0x22,
	0xe8, 0x3b, 0x11, 0x23, 0xb4, 0xc1, 0x5f, 0x3c, 0x50, 0xde, 0x08, 0xad, 0x61, 0xfb, 0x77, 0x07,
	0x6c, 0x9d, 0xf2, 0x78, 0x98, 0xe1, 0x40, 0xe0, 0x77, 0x19, 0x8b, 0xb3, 0x20, 0xd1, 0xba, 0x60,
	0x0d, 0x61, 0x1e, 0x65, 0x24, 0x15, 0x84, 0x51, 0x5d, 0xed, 0xaa, 0xbd, 0x55, 0x6f, 0xbe, 0xa4,
	0x7d, 0x00, 0x9b, 0x11, 0x4b, 0x12, 0xc2, 0x39, 0x61, 0xd4, 0xcf, 0x02, 0x81, 0xf5, 0xa5, 0x8a,
	0xe5, 0xc2, 0x9b, 0xd2, 0x52, 0xbe, 0x95, 0xd6, 0x5e, 0x4c, 0xc4, 0x59, 0x1e, 0xc2, 0x88, 0x25,
	0xcd, 0x51, 0x9a, 0xcf, 0x01, 0x47, 0xe7, 0x8e, 0x28, 0x52, 0xcc, 0xe1, 0x6b, 0x1c, 0x79, 0x1b,
	0x77, 0x63, 0xbc, 0x40, 0x60, 0x2d, 0x03, 0xdb, 0x3c, 0x0f, 0xdb, 0xc1, 0x98, 0x22, 0xbf, 0xb2,
	0x49, 0xef, 0x74, 0xd5, 0xde, 0xda, 0xc0, 0x80, 0xd2, 0x07, 0xd8, 0xfa, 0x00, 0xdf, 0xb7, 0x1e,
	0xba, 0x7b, 0xd5, 0xc6, 0xb3, 0xd2, 0x32, 0x8a, 0x20, 0xb9, 0x38, 0xb4, 0xff, 0x31, 0xc4, 0xbe,
	0xfa, 0x6e, 0xa9, 0xde, 0xd3, 0x3b, 0xe4, 0x98, 0xa2, 0xaa, 0x5f, 0x1b, 0x82, 0xcd, 0xa8, 0x3a,
	0x3f, 0xcb, 0xfc, 0x00, 0xa1, 0x0c, 0x73, 0xae, 0x2f, 0xd7, 0x87, 0x31, 0x66, 0xa5, 0xf5, 0x5c,
	0xce, 0x5b, 0x20, 0xd8, 0xde, 0x46, 0x53, 0x39, 0x92, 0x05, 0x2d, 0x06, 0x1b, 0x98, 0x46, 0x59,
	0x51, 0xfb, 0xe3, 0x9f, 0xe3, 0x42, 0x7f, 0x54, 0x6b, 0xde, 0x79, 0xa0, 0xf9, 0x88, 0x16, 0xee,
	0xfe, 0xac, 0xb4, 0x9e, 0xc9, 0xc9, 0xf7, 0xbb, 0xec, 0x2f, 0x9f, 0x0f, 0xd6, 0x8f, 0xff, 0x96,
	0x4e, 0x70, 0xe1, 0xad, 0xe3, 0xf9, 0xa5, 0xf6, 0x0a, 0x3c, 0x46, 0x38, 0x65, 0x9c, 0x08, 0x7d,
	0xa5, 0xdb, 0xe9, 0xad, 0x0d, 0x76, 0x61, 0x93, 0x95, 0x2a, 0x02, 0xb0, 0x89, 0x00, 0x1c, 0x32,
	0x42, 0xdd, 0xe5, 0xca, 0x14, 0xaf, 0xe5, 0x1f, 0x3e, 0xf9, 0x74, 0x6d, 0x29, 0x3f, 0xaf, 0x2d,
	0xc5, 0x7e, 0x0b, 0xf4, 0xc5, 0x5b, 0xf7, 0x30, 0x4f, 0x19, 0xe5, 0x58, 0x1b, 0x00, 0x90, 0xca,
	0x92, 0x4f, 0x50, 0x7d, 0xf9, 0xcb, 0xee, 0xf6, 0xaf, 0xd2, 0x5a, 0x22, 0x68, 0x56, 0x5a, 0xab,
	0x52, 0x35, 0x41, 0xb6, 0xb7, 0xda, 0xd0, 0xde, 0xa0, 0xc1, 0x08, 0x74, 0x4e, 0x79, 0xac, 0xf9,
	0x60, 0xfd, 0x7e, 0x92, 0x6c, 0xb8, 0xf8, 0x3a, 0xe0, 0xe2, 0xbe, 0xc6, 0xfe, 0xff, 0x39, 0xad,
	0x36, 0xf7, 0xe4, 0x66, 0x62, 0xaa, 0xb7, 0x13, 0x53, 0xfd, 0x31, 0x31, 0xd5, 0xab, 0xa9, 0xa9,
	0xdc, 0x4e, 0x4d, 0xe5, 0xeb, 0xd4, 0x54, 0x3e, 0xf6, 0xe7, 0x02, 0x27, 0xe7, 0x8d, 0x58, 0x4e,
	0x51, 0x50, 0xd9, 0xd6, 0x14, 0x9c, 0xcb, 0xf6, 0x15, 0xd4, 0xf9, 0x0b, 0x57, 0xea, 0x2b, 0x79,
	0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x64, 0x37, 0x9d, 0x88, 0xdd, 0x03, 0x00, 0x00,
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
	// CreateProgram defines a method for creating a new program.
	CreateProgram(ctx context.Context, in *MsgCreateProgram, opts ...grpc.CallOption) (*MsgCreateProgramResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateProgram(ctx context.Context, in *MsgCreateProgram, opts ...grpc.CallOption) (*MsgCreateProgramResponse, error) {
	out := new(MsgCreateProgramResponse)
	err := c.cc.Invoke(ctx, "/shentu.bounty.v1.Msg/CreateProgram", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateProgram defines a method for creating a new program.
	CreateProgram(context.Context, *MsgCreateProgram) (*MsgCreateProgramResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateProgram(ctx context.Context, req *MsgCreateProgram) (*MsgCreateProgramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProgram not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateProgram_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateProgram)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateProgram(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.bounty.v1.Msg/CreateProgram",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateProgram(ctx, req.(*MsgCreateProgram))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shentu.bounty.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProgram",
			Handler:    _Msg_CreateProgram_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shentu/bounty/v1/tx.proto",
}

func (m *MsgCreateProgram) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateProgram) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateProgram) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deposit) > 0 {
		for iNdEx := len(m.Deposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.EncryptionKey != nil {
		{
			size, err := m.EncryptionKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CreatorAddress) > 0 {
		i -= len(m.CreatorAddress)
		copy(dAtA[i:], m.CreatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CreatorAddress)))
		i--
		dAtA[i] = 0x22
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.SubmissionEndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.SubmissionEndTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTx(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	{
		size := m.CommissionRate.Size()
		i -= size
		if _, err := m.CommissionRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateProgramResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateProgramResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateProgramResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProgramId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ProgramId))
		i--
		dAtA[i] = 0x8
	}
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
func (m *MsgCreateProgram) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.CommissionRate.Size()
	n += 1 + l + sovTx(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.SubmissionEndTime)
	n += 1 + l + sovTx(uint64(l))
	l = len(m.CreatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.EncryptionKey != nil {
		l = m.EncryptionKey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Deposit) > 0 {
		for _, e := range m.Deposit {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreateProgramResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProgramId != 0 {
		n += 1 + sovTx(uint64(m.ProgramId))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateProgram) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateProgram: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateProgram: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommissionRate", wireType)
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
			if err := m.CommissionRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmissionEndTime", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.SubmissionEndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatorAddress", wireType)
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
			m.CreatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptionKey", wireType)
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
			if m.EncryptionKey == nil {
				m.EncryptionKey = &types.Any{}
			}
			if err := m.EncryptionKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposit", wireType)
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
			m.Deposit = append(m.Deposit, types1.Coin{})
			if err := m.Deposit[len(m.Deposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
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
func (m *MsgCreateProgramResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateProgramResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateProgramResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProgramId", wireType)
			}
			m.ProgramId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProgramId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
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
