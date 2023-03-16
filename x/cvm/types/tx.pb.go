// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/cvm/v1alpha1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	payload "github.com/hyperledger/burrow/txs/payload"
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

type MsgCall struct {
	Caller string `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty" yaml:"caller"`
	Callee string `protobuf:"bytes,2,opt,name=callee,proto3" json:"callee,omitempty" yaml:"callee"`
	Value  uint64 `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty" yaml:"value"`
	Data   []byte `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty" yaml:"data"`
}

func (m *MsgCall) Reset()         { *m = MsgCall{} }
func (m *MsgCall) String() string { return proto.CompactTextString(m) }
func (*MsgCall) ProtoMessage()    {}
func (*MsgCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f9b3577718d7e8, []int{0}
}
func (m *MsgCall) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCall.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCall.Merge(m, src)
}
func (m *MsgCall) XXX_Size() int {
	return m.Size()
}
func (m *MsgCall) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCall.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCall proto.InternalMessageInfo

func (m *MsgCall) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

func (m *MsgCall) GetCallee() string {
	if m != nil {
		return m.Callee
	}
	return ""
}

func (m *MsgCall) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MsgCall) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type MsgCallResponse struct {
	Result []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty" yaml:"result"`
}

func (m *MsgCallResponse) Reset()         { *m = MsgCallResponse{} }
func (m *MsgCallResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCallResponse) ProtoMessage()    {}
func (*MsgCallResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f9b3577718d7e8, []int{1}
}
func (m *MsgCallResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCallResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCallResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCallResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCallResponse.Merge(m, src)
}
func (m *MsgCallResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCallResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCallResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCallResponse proto.InternalMessageInfo

func (m *MsgCallResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

type MsgDeploy struct {
	// Caller is the sender of the CVM-message.
	Caller string `protobuf:"bytes,1,opt,name=caller,proto3" json:"caller,omitempty" yaml:"caller"`
	// Value is the amount of CTK transferred with the call.
	Value uint64 `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty" yaml:"value"`
	// Code is the contract byte code.
	Code []byte `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty" yaml:"code"`
	// Abi is the Solidity ABI bytes for the contract code.
	Abi string `protobuf:"bytes,4,opt,name=abi,proto3" json:"abi,omitempty" yaml:"abi"`
	// Meta is the metadata for the contract.
	Meta []*payload.ContractMeta `protobuf:"bytes,5,rep,name=meta,proto3" json:"meta,omitempty" yaml:"meta"`
	// is_eWASM is true if the code is EWASM code.
	IsEWASM bool `protobuf:"varint,6,opt,name=is_eWASM,json=isEWASM,proto3" json:"is_eWASM,omitempty" yaml:"is_EWASM"`
	// is_runtime is true if the code is runtime code.
	IsRuntime bool `protobuf:"varint,7,opt,name=is_runtime,json=isRuntime,proto3" json:"is_runtime,omitempty" yaml:"is_runtime"`
}

func (m *MsgDeploy) Reset()         { *m = MsgDeploy{} }
func (m *MsgDeploy) String() string { return proto.CompactTextString(m) }
func (*MsgDeploy) ProtoMessage()    {}
func (*MsgDeploy) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f9b3577718d7e8, []int{2}
}
func (m *MsgDeploy) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeploy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeploy.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeploy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeploy.Merge(m, src)
}
func (m *MsgDeploy) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeploy) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeploy.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeploy proto.InternalMessageInfo

func (m *MsgDeploy) GetCaller() string {
	if m != nil {
		return m.Caller
	}
	return ""
}

func (m *MsgDeploy) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MsgDeploy) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *MsgDeploy) GetAbi() string {
	if m != nil {
		return m.Abi
	}
	return ""
}

func (m *MsgDeploy) GetMeta() []*payload.ContractMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *MsgDeploy) GetIsEWASM() bool {
	if m != nil {
		return m.IsEWASM
	}
	return false
}

func (m *MsgDeploy) GetIsRuntime() bool {
	if m != nil {
		return m.IsRuntime
	}
	return false
}

type MsgDeployResponse struct {
	Result []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty" yaml:"result"`
}

func (m *MsgDeployResponse) Reset()         { *m = MsgDeployResponse{} }
func (m *MsgDeployResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeployResponse) ProtoMessage()    {}
func (*MsgDeployResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_83f9b3577718d7e8, []int{3}
}
func (m *MsgDeployResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeployResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeployResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeployResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeployResponse.Merge(m, src)
}
func (m *MsgDeployResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeployResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeployResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeployResponse proto.InternalMessageInfo

func (m *MsgDeployResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgCall)(nil), "shentu.cvm.v1alpha1.MsgCall")
	proto.RegisterType((*MsgCallResponse)(nil), "shentu.cvm.v1alpha1.MsgCallResponse")
	proto.RegisterType((*MsgDeploy)(nil), "shentu.cvm.v1alpha1.MsgDeploy")
	proto.RegisterType((*MsgDeployResponse)(nil), "shentu.cvm.v1alpha1.MsgDeployResponse")
}

func init() { proto.RegisterFile("shentu/cvm/v1alpha1/tx.proto", fileDescriptor_83f9b3577718d7e8) }

var fileDescriptor_83f9b3577718d7e8 = []byte{
	// 521 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x8a, 0xd3, 0x40,
	0x18, 0xc7, 0x9b, 0xb6, 0xdb, 0x6e, 0x67, 0xab, 0xb5, 0xd9, 0x5d, 0x08, 0x65, 0x4d, 0xc2, 0x28,
	0x4b, 0xbd, 0x24, 0xec, 0xea, 0x69, 0x11, 0xc1, 0xac, 0x82, 0x08, 0x05, 0x19, 0x0f, 0x82, 0x97,
	0x32, 0x4d, 0xc7, 0x34, 0x30, 0xc9, 0x84, 0xcc, 0xa4, 0x6e, 0xdf, 0xc2, 0x77, 0xf0, 0xea, 0x7b,
	0xe8, 0x71, 0x8f, 0x9e, 0x82, 0xb4, 0x6f, 0x90, 0x27, 0x90, 0xcc, 0xa4, 0x61, 0x85, 0xba, 0xe2,
	0xde, 0x86, 0xff, 0xef, 0xff, 0x7d, 0xf9, 0xfe, 0xf9, 0x66, 0xc0, 0x09, 0x5f, 0x90, 0x58, 0x64,
	0xae, 0xbf, 0x8c, 0xdc, 0xe5, 0x19, 0xa6, 0xc9, 0x02, 0x9f, 0xb9, 0xe2, 0xca, 0x49, 0x52, 0x26,
	0x98, 0x7e, 0xa8, 0xa8, 0xe3, 0x2f, 0x23, 0x67, 0x4b, 0x47, 0x47, 0x01, 0x0b, 0x98, 0xe4, 0x6e,
	0x79, 0x52, 0xd6, 0xd1, 0xc3, 0x5d, 0x8d, 0xca, 0x3a, 0x85, 0x8f, 0x66, 0x59, 0x9a, 0xb2, 0xcf,
	0x6e, 0x82, 0x57, 0x94, 0xe1, 0xb9, 0x52, 0xe1, 0x37, 0x0d, 0x74, 0x27, 0x3c, 0xb8, 0xc4, 0x94,
	0xea, 0x4f, 0x40, 0xc7, 0xc7, 0x94, 0x92, 0xd4, 0xd0, 0x6c, 0x6d, 0xdc, 0xf3, 0x86, 0x45, 0x6e,
	0xdd, 0x5b, 0xe1, 0x88, 0x5e, 0x40, 0xa5, 0x43, 0x54, 0x19, 0x6a, 0x2b, 0x31, 0x9a, 0x3b, 0xad,
	0x64, 0x6b, 0x25, 0xfa, 0x29, 0xd8, 0x5b, 0x62, 0x9a, 0x11, 0xa3, 0x65, 0x6b, 0xe3, 0xb6, 0xf7,
	0xa0, 0xc8, 0xad, 0xbe, 0x72, 0x4a, 0x19, 0x22, 0x85, 0xf5, 0x47, 0xa0, 0x3d, 0xc7, 0x02, 0x1b,
	0x6d, 0x5b, 0x1b, 0xf7, 0xbd, 0x41, 0x91, 0x5b, 0x07, 0xca, 0x56, 0xaa, 0x10, 0x49, 0x08, 0x9f,
	0x83, 0x41, 0x35, 0x2d, 0x22, 0x3c, 0x61, 0x31, 0x27, 0xe5, 0x28, 0x29, 0xe1, 0x19, 0x15, 0x72,
	0xea, 0xfe, 0xcd, 0x51, 0x94, 0x0e, 0x51, 0x65, 0x80, 0xdf, 0x9b, 0xa0, 0x37, 0xe1, 0xc1, 0x2b,
	0x92, 0x50, 0xb6, 0xfa, 0x9f, 0xb8, 0x75, 0x86, 0xe6, 0x3f, 0x33, 0xf8, 0x6c, 0xae, 0xa2, 0xfe,
	0x91, 0xa1, 0x54, 0x21, 0x92, 0x50, 0xb7, 0x41, 0x0b, 0xcf, 0x42, 0x99, 0xb3, 0xe7, 0xdd, 0x2f,
	0x72, 0x0b, 0x28, 0x0f, 0x9e, 0x85, 0x10, 0x95, 0x48, 0xbf, 0x00, 0xed, 0x88, 0x08, 0x6c, 0xec,
	0xd9, 0xad, 0xf1, 0xc1, 0xf9, 0xb1, 0xb3, 0x5d, 0xd9, 0x25, 0x8b, 0x45, 0x8a, 0x7d, 0x31, 0x21,
	0x02, 0xdf, 0xec, 0x5e, 0x9a, 0x21, 0x92, 0x35, 0xba, 0x03, 0xf6, 0x43, 0x3e, 0x25, 0x1f, 0x5e,
	0xbe, 0x9f, 0x18, 0x1d, 0x5b, 0x1b, 0xef, 0x7b, 0x87, 0x45, 0x6e, 0x0d, 0x94, 0x31, 0xe4, 0xd3,
	0xd7, 0x25, 0x81, 0xa8, 0x1b, 0x72, 0x79, 0xd2, 0x9f, 0x01, 0x10, 0xf2, 0x69, 0x9a, 0xc5, 0x22,
	0x8c, 0x88, 0xd1, 0x95, 0x15, 0xc7, 0x45, 0x6e, 0x0d, 0xeb, 0x8a, 0x8a, 0x41, 0xd4, 0x0b, 0x39,
	0xaa, 0xce, 0x2f, 0xc0, 0xb0, 0xfe, 0x91, 0x77, 0xd8, 0xc4, 0xf9, 0x57, 0x0d, 0xb4, 0x26, 0x3c,
	0xd0, 0xdf, 0x82, 0xb6, 0xbc, 0x7a, 0x27, 0xce, 0x8e, 0x7b, 0xee, 0x54, 0xab, 0x1e, 0x3d, 0xbe,
	0x8d, 0xd6, 0x9f, 0x7f, 0x07, 0x3a, 0xd5, 0x66, 0xcd, 0xbf, 0xf9, 0x15, 0x1f, 0x9d, 0xde, 0xce,
	0xb7, 0x1d, 0xbd, 0x37, 0x3f, 0xd6, 0xa6, 0x76, 0xbd, 0x36, 0xb5, 0x5f, 0x6b, 0x53, 0xfb, 0xb2,
	0x31, 0x1b, 0xd7, 0x1b, 0xb3, 0xf1, 0x73, 0x63, 0x36, 0x3e, 0x3a, 0x41, 0x28, 0x16, 0xd9, 0xcc,
	0xf1, 0x59, 0xe4, 0xaa, 0x5e, 0x9f, 0x58, 0x16, 0xcf, 0xb1, 0x08, 0x59, 0x5c, 0x09, 0xee, 0x95,
	0x7c, 0x89, 0x62, 0x95, 0x10, 0x3e, 0xeb, 0xc8, 0xd7, 0xf6, 0xf4, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x59, 0x75, 0x94, 0x4d, 0xed, 0x03, 0x00, 0x00,
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
	Call(ctx context.Context, in *MsgCall, opts ...grpc.CallOption) (*MsgCallResponse, error)
	Deploy(ctx context.Context, in *MsgDeploy, opts ...grpc.CallOption) (*MsgDeployResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Call(ctx context.Context, in *MsgCall, opts ...grpc.CallOption) (*MsgCallResponse, error) {
	out := new(MsgCallResponse)
	err := c.cc.Invoke(ctx, "/shentu.cvm.v1alpha1.Msg/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Deploy(ctx context.Context, in *MsgDeploy, opts ...grpc.CallOption) (*MsgDeployResponse, error) {
	out := new(MsgDeployResponse)
	err := c.cc.Invoke(ctx, "/shentu.cvm.v1alpha1.Msg/Deploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Call(context.Context, *MsgCall) (*MsgCallResponse, error)
	Deploy(context.Context, *MsgDeploy) (*MsgDeployResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Call(ctx context.Context, req *MsgCall) (*MsgCallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}
func (*UnimplementedMsgServer) Deploy(ctx context.Context, req *MsgDeploy) (*MsgDeployResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCall)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.cvm.v1alpha1.Msg/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Call(ctx, req.(*MsgCall))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeploy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.cvm.v1alpha1.Msg/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Deploy(ctx, req.(*MsgDeploy))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shentu.cvm.v1alpha1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Msg_Call_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Msg_Deploy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shentu/cvm/v1alpha1/tx.proto",
}

func (m *MsgCall) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCall) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCall) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if m.Value != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Callee) > 0 {
		i -= len(m.Callee)
		copy(dAtA[i:], m.Callee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Callee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Caller) > 0 {
		i -= len(m.Caller)
		copy(dAtA[i:], m.Caller)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Caller)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCallResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCallResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCallResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeploy) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeploy) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeploy) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsRuntime {
		i--
		if m.IsRuntime {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.IsEWASM {
		i--
		if m.IsEWASM {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.Meta) > 0 {
		for iNdEx := len(m.Meta) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Meta[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Abi) > 0 {
		i -= len(m.Abi)
		copy(dAtA[i:], m.Abi)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Abi)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Code) > 0 {
		i -= len(m.Code)
		copy(dAtA[i:], m.Code)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Code)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Value != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Caller) > 0 {
		i -= len(m.Caller)
		copy(dAtA[i:], m.Caller)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Caller)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeployResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeployResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeployResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Result) > 0 {
		i -= len(m.Result)
		copy(dAtA[i:], m.Result)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Result)))
		i--
		dAtA[i] = 0xa
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
func (m *MsgCall) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Caller)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Callee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovTx(uint64(m.Value))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCallResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeploy) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Caller)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovTx(uint64(m.Value))
	}
	l = len(m.Code)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Abi)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Meta) > 0 {
		for _, e := range m.Meta {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if m.IsEWASM {
		n += 2
	}
	if m.IsRuntime {
		n += 2
	}
	return n
}

func (m *MsgDeployResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Result)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCall) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCall: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCall: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Caller", wireType)
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
			m.Caller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Callee", wireType)
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
			m.Callee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
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
func (m *MsgCallResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCallResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCallResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
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
func (m *MsgDeploy) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeploy: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeploy: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Caller", wireType)
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
			m.Caller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Code = append(m.Code[:0], dAtA[iNdEx:postIndex]...)
			if m.Code == nil {
				m.Code = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Abi", wireType)
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
			m.Abi = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
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
			m.Meta = append(m.Meta, &payload.ContractMeta{})
			if err := m.Meta[len(m.Meta)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsEWASM", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsEWASM = bool(v != 0)
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsRuntime", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsRuntime = bool(v != 0)
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
func (m *MsgDeployResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeployResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeployResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Result = append(m.Result[:0], dAtA[iNdEx:postIndex]...)
			if m.Result == nil {
				m.Result = []byte{}
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
