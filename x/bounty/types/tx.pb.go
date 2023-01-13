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

type MsgSubmitFinding struct {
	Title            string        `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Desc             string        `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty" yaml:"description"`
	ProgramId        uint64        `protobuf:"varint,3,opt,name=program_id,json=programId,proto3" json:"program_id,omitempty" yaml:"program_id"`
	SeverityLevel    SeverityLevel `protobuf:"varint,4,opt,name=severity_level,json=severityLevel,proto3,enum=shentu.bounty.v1.SeverityLevel" json:"severity_level,omitempty" yaml:"severity_level"`
	Poc              string        `protobuf:"bytes,5,opt,name=poc,proto3" json:"poc,omitempty" yaml:"poc"`
	SubmitterAddress string        `protobuf:"bytes,6,opt,name=submitter_address,json=submitterAddress,proto3" json:"submitter_address,omitempty" yaml:"submitter_address"`
}

func (m *MsgSubmitFinding) Reset()         { *m = MsgSubmitFinding{} }
func (m *MsgSubmitFinding) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitFinding) ProtoMessage()    {}
func (*MsgSubmitFinding) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4b4296bac3db30, []int{2}
}
func (m *MsgSubmitFinding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitFinding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitFinding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitFinding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitFinding.Merge(m, src)
}
func (m *MsgSubmitFinding) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitFinding) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitFinding.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitFinding proto.InternalMessageInfo

type MsgSubmitFindingResponse struct {
	FindingId uint64 `protobuf:"varint,1,opt,name=finding_id,json=findingId,proto3" json:"finding_id" yaml:"finding_id"`
}

func (m *MsgSubmitFindingResponse) Reset()         { *m = MsgSubmitFindingResponse{} }
func (m *MsgSubmitFindingResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitFindingResponse) ProtoMessage()    {}
func (*MsgSubmitFindingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4b4296bac3db30, []int{3}
}
func (m *MsgSubmitFindingResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitFindingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitFindingResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitFindingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitFindingResponse.Merge(m, src)
}
func (m *MsgSubmitFindingResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitFindingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitFindingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitFindingResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateProgram)(nil), "shentu.bounty.v1.MsgCreateProgram")
	proto.RegisterType((*MsgCreateProgramResponse)(nil), "shentu.bounty.v1.MsgCreateProgramResponse")
	proto.RegisterType((*MsgSubmitFinding)(nil), "shentu.bounty.v1.MsgSubmitFinding")
	proto.RegisterType((*MsgSubmitFindingResponse)(nil), "shentu.bounty.v1.MsgSubmitFindingResponse")
}

func init() { proto.RegisterFile("shentu/bounty/v1/tx.proto", fileDescriptor_1e4b4296bac3db30) }

var fileDescriptor_1e4b4296bac3db30 = []byte{
	// 775 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x6a, 0xeb, 0x46,
	0x14, 0xb5, 0x62, 0xbf, 0xb4, 0x9e, 0x60, 0xbf, 0xbc, 0x79, 0x79, 0x0f, 0xd9, 0xb4, 0x1e, 0xa3,
	0xc2, 0x23, 0x04, 0x22, 0x61, 0xb7, 0x9b, 0x66, 0x17, 0xa5, 0x29, 0x84, 0x34, 0xa5, 0x28, 0x85,
	0x42, 0x37, 0x42, 0x96, 0x26, 0xca, 0x10, 0x4b, 0x23, 0x34, 0x63, 0x13, 0xfd, 0x41, 0x97, 0xf9,
	0x84, 0x7c, 0x44, 0xff, 0xa1, 0xa1, 0xab, 0x2c, 0x4b, 0x17, 0x6a, 0x49, 0xba, 0x28, 0x59, 0x8a,
	0x7e, 0x40, 0xd1, 0xcc, 0xd8, 0x92, 0x9d, 0x94, 0xae, 0xec, 0xb9, 0xe7, 0xcc, 0x9d, 0x33, 0xe7,
	0x1e, 0x0d, 0xe8, 0xb1, 0x4b, 0x1c, 0xf3, 0x99, 0x35, 0xa1, 0xb3, 0x98, 0x67, 0xd6, 0x7c, 0x64,
	0xf1, 0x6b, 0x33, 0x49, 0x29, 0xa7, 0x70, 0x5b, 0x42, 0xa6, 0x84, 0xcc, 0xf9, 0xa8, 0xbf, 0x13,
	0xd2, 0x90, 0x0a, 0xd0, 0x2a, 0xff, 0x49, 0x5e, 0x1f, 0x85, 0x94, 0x86, 0x53, 0x6c, 0x89, 0xd5,
	0x64, 0x76, 0x61, 0x71, 0x12, 0x61, 0xc6, 0xbd, 0x28, 0x51, 0x84, 0x9e, 0x4f, 0x59, 0x44, 0x99,
	0x2b, 0x77, 0xca, 0xc5, 0x02, 0x5a, 0xdf, 0xeb, 0xc5, 0x99, 0x82, 0x06, 0x92, 0x68, 0x4d, 0x3c,
	0x86, 0xad, 0xf9, 0x68, 0x82, 0xb9, 0x37, 0xb2, 0x7c, 0x4a, 0x62, 0x85, 0x7f, 0xfa, 0x4c, 0xb9,
	0x12, 0x2a, 0x60, 0xe3, 0x9f, 0x26, 0xd8, 0x3e, 0x63, 0xe1, 0x51, 0x8a, 0x3d, 0x8e, 0xbf, 0x4b,
	0x69, 0x98, 0x7a, 0x11, 0x1c, 0x82, 0xad, 0x00, 0x33, 0x3f, 0x25, 0x09, 0x27, 0x34, 0xd6, 0xb5,
	0xa1, 0xb6, 0xdb, 0x76, 0xea, 0x25, 0xf8, 0x03, 0x78, 0xed, 0xd3, 0x28, 0x22, 0x8c, 0x11, 0x1a,
	0xbb, 0xa9, 0xc7, 0xb1, 0xbe, 0x51, 0xb2, 0x6c, 0xf3, 0x2e, 0x47, 0x8d, 0xdf, 0x73, 0xf4, 0x21,
	0x24, 0xfc, 0x72, 0x36, 0x31, 0x7d, 0x1a, 0xa9, 0xab, 0xa8, 0x9f, 0x7d, 0x16, 0x5c, 0x59, 0x3c,
	0x4b, 0x30, 0x33, 0xbf, 0xc2, 0xbe, 0xd3, 0xad, 0xda, 0x38, 0x1e, 0xc7, 0x30, 0x05, 0x6f, 0xd9,
	0x6c, 0xb2, 0x68, 0x8c, 0xe3, 0xc0, 0x2d, 0x6d, 0xd2, 0x9b, 0x43, 0x6d, 0x77, 0x6b, 0xdc, 0x37,
	0xa5, 0x0f, 0xe6, 0xc2, 0x07, 0xf3, 0xfb, 0x85, 0x87, 0xf6, 0x87, 0xf2, 0xe0, 0x22, 0x47, 0xfd,
	0xcc, 0x8b, 0xa6, 0x07, 0xc6, 0x0b, 0x4d, 0x8c, 0x9b, 0x3f, 0x90, 0xe6, 0xbc, 0xa9, 0x90, 0xe3,
	0x38, 0x28, 0xf7, 0xc3, 0x23, 0xf0, 0xda, 0x2f, 0xef, 0x4f, 0x53, 0xd7, 0x0b, 0x82, 0x14, 0x33,
	0xa6, 0xb7, 0xc4, 0x65, 0xfa, 0x45, 0x8e, 0xde, 0xcb, 0x7e, 0x6b, 0x04, 0xc3, 0xe9, 0xaa, 0xca,
	0xa1, 0x2c, 0xc0, 0x10, 0x74, 0x71, 0xec, 0xa7, 0x99, 0xf0, 0xc7, 0xbd, 0xc2, 0x99, 0xfe, 0x4a,
	0x68, 0xde, 0x79, 0xa6, 0xf9, 0x30, 0xce, 0xec, 0xbd, 0x22, 0x47, 0xef, 0x64, 0xe7, 0xd5, 0x5d,
	0xc6, 0xaf, 0x3f, 0xef, 0x77, 0x8e, 0x97, 0xa5, 0x53, 0x9c, 0x39, 0x1d, 0x5c, 0x5f, 0xc2, 0x2f,
	0xc1, 0x47, 0x01, 0x4e, 0x28, 0x23, 0x5c, 0xdf, 0x1c, 0x36, 0x77, 0xb7, 0xc6, 0x3d, 0x53, 0x65,
	0xa5, 0x8c, 0x80, 0xa9, 0x22, 0x60, 0x1e, 0x51, 0x12, 0xdb, 0xad, 0xd2, 0x14, 0x67, 0xc1, 0x3f,
	0xf8, 0xf8, 0xa7, 0x5b, 0xd4, 0xf8, 0xfb, 0x16, 0x35, 0x8c, 0x6f, 0x81, 0xbe, 0x3e, 0x75, 0x07,
	0xb3, 0x84, 0xc6, 0x0c, 0xc3, 0x31, 0x00, 0x89, 0x2c, 0xb9, 0x24, 0x10, 0xc3, 0x6f, 0xd9, 0x6f,
	0x9f, 0x72, 0xb4, 0x41, 0x82, 0x22, 0x47, 0x6d, 0xa9, 0x9a, 0x04, 0x86, 0xd3, 0x56, 0xb4, 0x93,
	0xc0, 0xf8, 0x6b, 0x43, 0xc4, 0xe8, 0xbc, 0xf4, 0x96, 0x7f, 0x4d, 0xe2, 0x80, 0xc4, 0x21, 0xdc,
	0x01, 0xaf, 0x38, 0xe1, 0x53, 0xac, 0x02, 0x24, 0x17, 0x70, 0x0f, 0xb4, 0xca, 0x24, 0xa9, 0xbc,
	0xbc, 0x2f, 0x72, 0x04, 0x65, 0xcb, 0x5a, 0xbe, 0x0c, 0x47, 0x70, 0xe0, 0x17, 0x2b, 0x52, 0x9a,
	0x42, 0xca, 0xbb, 0x22, 0x47, 0x6f, 0xe4, 0x8e, 0x0a, 0xab, 0x8b, 0x81, 0x1e, 0xe8, 0x32, 0x3c,
	0xc7, 0x29, 0xe1, 0x99, 0x3b, 0xc5, 0x73, 0x3c, 0x15, 0xe3, 0xec, 0x8e, 0x91, 0xb9, 0xfe, 0xa9,
	0x9a, 0xe7, 0x8a, 0xf7, 0x4d, 0x49, 0xb3, 0x7b, 0xd5, 0x54, 0x56, 0x1b, 0x18, 0x4e, 0x87, 0xd5,
	0x99, 0x70, 0x08, 0x9a, 0x09, 0xf5, 0xc5, 0x88, 0xdb, 0x76, 0xb7, 0xc8, 0x11, 0x50, 0x8a, 0xa8,
	0x6f, 0x38, 0x25, 0x04, 0x4f, 0x80, 0x4c, 0x1a, 0xe7, 0xb8, 0x8a, 0xd5, 0xa6, 0xe0, 0x7f, 0x52,
	0xe4, 0x48, 0xaf, 0xc5, 0xb4, 0x4e, 0x31, 0x9c, 0xed, 0x65, 0x4d, 0x45, 0xab, 0x36, 0xb6, 0x40,
	0x8c, 0x6d, 0xc5, 0xe5, 0xe5, 0xd8, 0x6c, 0x00, 0x2e, 0x64, 0xa9, 0x1a, 0xdb, 0x67, 0x4f, 0x39,
	0xaa, 0x55, 0x2b, 0xe7, 0xaa, 0x9a, 0xe1, 0xb4, 0xd5, 0xe2, 0x24, 0x38, 0x68, 0x95, 0x27, 0x8d,
	0x7f, 0xd1, 0x40, 0xf3, 0x8c, 0x85, 0xd0, 0x05, 0x9d, 0xd5, 0x77, 0xc1, 0x78, 0x6e, 0xe0, 0x7a,
	0x8a, 0xfa, 0x7b, 0xff, 0xcf, 0x59, 0x4a, 0x76, 0x41, 0x67, 0x35, 0x31, 0x2f, 0x1f, 0xb0, 0xc2,
	0xf9, 0x8f, 0x03, 0x5e, 0xf4, 0xc4, 0x3e, 0xbd, 0x7b, 0x18, 0x68, 0xf7, 0x0f, 0x03, 0xed, 0xcf,
	0x87, 0x81, 0x76, 0xf3, 0x38, 0x68, 0xdc, 0x3f, 0x0e, 0x1a, 0xbf, 0x3d, 0x0e, 0x1a, 0x3f, 0x8e,
	0x6a, 0xef, 0x93, 0xec, 0x77, 0x41, 0x67, 0x71, 0xe0, 0x95, 0xe9, 0x53, 0x05, 0xeb, 0x7a, 0xf1,
	0x68, 0x8a, 0xe7, 0x6a, 0xb2, 0x29, 0xbe, 0xe0, 0xcf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x67,
	0x6e, 0x84, 0x21, 0x0c, 0x06, 0x00, 0x00,
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
	// SubmitFinding defines a method for submitting a new finding.
	SubmitFinding(ctx context.Context, in *MsgSubmitFinding, opts ...grpc.CallOption) (*MsgSubmitFindingResponse, error)
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

func (c *msgClient) SubmitFinding(ctx context.Context, in *MsgSubmitFinding, opts ...grpc.CallOption) (*MsgSubmitFindingResponse, error) {
	out := new(MsgSubmitFindingResponse)
	err := c.cc.Invoke(ctx, "/shentu.bounty.v1.Msg/SubmitFinding", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateProgram defines a method for creating a new program.
	CreateProgram(context.Context, *MsgCreateProgram) (*MsgCreateProgramResponse, error)
	// SubmitFinding defines a method for submitting a new finding.
	SubmitFinding(context.Context, *MsgSubmitFinding) (*MsgSubmitFindingResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateProgram(ctx context.Context, req *MsgCreateProgram) (*MsgCreateProgramResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProgram not implemented")
}
func (*UnimplementedMsgServer) SubmitFinding(ctx context.Context, req *MsgSubmitFinding) (*MsgSubmitFindingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitFinding not implemented")
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

func _Msg_SubmitFinding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitFinding)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitFinding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.bounty.v1.Msg/SubmitFinding",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitFinding(ctx, req.(*MsgSubmitFinding))
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
		{
			MethodName: "SubmitFinding",
			Handler:    _Msg_SubmitFinding_Handler,
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

func (m *MsgSubmitFinding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitFinding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitFinding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SubmitterAddress) > 0 {
		i -= len(m.SubmitterAddress)
		copy(dAtA[i:], m.SubmitterAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SubmitterAddress)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Poc) > 0 {
		i -= len(m.Poc)
		copy(dAtA[i:], m.Poc)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Poc)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SeverityLevel != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.SeverityLevel))
		i--
		dAtA[i] = 0x20
	}
	if m.ProgramId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ProgramId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Desc) > 0 {
		i -= len(m.Desc)
		copy(dAtA[i:], m.Desc)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Desc)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitFindingResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitFindingResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitFindingResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FindingId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.FindingId))
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

func (m *MsgSubmitFinding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Desc)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.ProgramId != 0 {
		n += 1 + sovTx(uint64(m.ProgramId))
	}
	if m.SeverityLevel != 0 {
		n += 1 + sovTx(uint64(m.SeverityLevel))
	}
	l = len(m.Poc)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SubmitterAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitFindingResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FindingId != 0 {
		n += 1 + sovTx(uint64(m.FindingId))
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
func (m *MsgSubmitFinding) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitFinding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitFinding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
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
			m.Desc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeverityLevel", wireType)
			}
			m.SeverityLevel = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SeverityLevel |= SeverityLevel(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Poc", wireType)
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
			m.Poc = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmitterAddress", wireType)
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
			m.SubmitterAddress = string(dAtA[iNdEx:postIndex])
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
func (m *MsgSubmitFindingResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitFindingResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitFindingResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FindingId", wireType)
			}
			m.FindingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FindingId |= uint64(b&0x7F) << shift
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
