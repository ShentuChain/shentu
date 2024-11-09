// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/gov/v1/genesis.proto

package v1

import (
	fmt "fmt"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the gov module's genesis state.
type GenesisState struct {
	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty"`
	// deposits defines all the deposits present at genesis.
	Deposits []*v1.Deposit `protobuf:"bytes,2,rep,name=deposits,proto3" json:"deposits,omitempty"`
	// votes defines all the votes present at genesis.
	Votes []*v1.Vote `protobuf:"bytes,3,rep,name=votes,proto3" json:"votes,omitempty"`
	// proposals defines all the proposals present at genesis.
	Proposals []*v1.Proposal `protobuf:"bytes,4,rep,name=proposals,proto3" json:"proposals,omitempty"`
	// params defines all the parameters of related to deposit.
	DepositParams *v1.DepositParams `protobuf:"bytes,5,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params,omitempty"`
	// params defines all the parameters of related to voting.
	VotingParams *v1.VotingParams `protobuf:"bytes,6,opt,name=voting_params,json=votingParams,proto3" json:"voting_params,omitempty"`
	// params defines all the parameters of related to tally.
	TallyParams *v1.TallyParams `protobuf:"bytes,7,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params,omitempty"`
	// params defines all the paramaters of x/gov module.
	//
	// Since: cosmos-sdk 0.47
	Params *v1.Params `protobuf:"bytes,8,opt,name=params,proto3" json:"params,omitempty"`
	// The constitution allows builders to lay a foundation and define purpose.
	// This is an immutable string set in genesis.
	// There are no amendments, to go outside of scope, just fork.
	// constitution is an immutable string in genesis for a chain builder to lay out their vision, ideas and ideals.
	//
	// Since: cosmos-sdk 0.50
	Constitution string `protobuf:"bytes,9,opt,name=constitution,proto3" json:"constitution,omitempty"`
	// params defines all the parameters of related to custom.
	CustomParams *CustomParams `protobuf:"bytes,10,opt,name=custom_params,json=customParams,proto3" json:"custom_params,omitempty"`
	// proposals that require and have passed cert votes.
	CertVotedProposalIds []uint64 `protobuf:"varint,11,rep,packed,name=cert_voted_proposal_ids,json=certVotedProposalIds,proto3" json:"cert_voted_proposal_ids,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_a3db44d59cab7379, []int{0}
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

func (m *GenesisState) GetStartingProposalId() uint64 {
	if m != nil {
		return m.StartingProposalId
	}
	return 0
}

func (m *GenesisState) GetDeposits() []*v1.Deposit {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetVotes() []*v1.Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisState) GetProposals() []*v1.Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisState) GetDepositParams() *v1.DepositParams {
	if m != nil {
		return m.DepositParams
	}
	return nil
}

func (m *GenesisState) GetVotingParams() *v1.VotingParams {
	if m != nil {
		return m.VotingParams
	}
	return nil
}

func (m *GenesisState) GetTallyParams() *v1.TallyParams {
	if m != nil {
		return m.TallyParams
	}
	return nil
}

func (m *GenesisState) GetParams() *v1.Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetConstitution() string {
	if m != nil {
		return m.Constitution
	}
	return ""
}

func (m *GenesisState) GetCustomParams() *CustomParams {
	if m != nil {
		return m.CustomParams
	}
	return nil
}

func (m *GenesisState) GetCertVotedProposalIds() []uint64 {
	if m != nil {
		return m.CertVotedProposalIds
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "shentu.gov.v1.GenesisState")
}

func init() { proto.RegisterFile("shentu/gov/v1/genesis.proto", fileDescriptor_a3db44d59cab7379) }

var fileDescriptor_a3db44d59cab7379 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x3f, 0x6f, 0xd3, 0x40,
	0x14, 0x8f, 0x49, 0x1a, 0x9a, 0x4b, 0xc2, 0x70, 0x04, 0x62, 0xa5, 0xc8, 0xb2, 0x3a, 0x99, 0x01,
	0x9b, 0x04, 0x75, 0x44, 0x42, 0x14, 0x09, 0x21, 0x96, 0xca, 0x20, 0x06, 0x96, 0xc8, 0xb5, 0x0f,
	0xd7, 0x52, 0xe2, 0x67, 0xf9, 0x3d, 0x9f, 0xe8, 0xb7, 0xe0, 0x63, 0x31, 0x76, 0x64, 0x44, 0xc9,
	0xc2, 0xc7, 0x40, 0xbe, 0xf3, 0x35, 0x8e, 0x61, 0xf3, 0xf3, 0xef, 0xcf, 0xfd, 0xde, 0xdd, 0x8f,
	0x9d, 0xe1, 0x8d, 0xc8, 0xa9, 0x0a, 0x52, 0x90, 0x81, 0x5c, 0x06, 0xa9, 0xc8, 0x05, 0x66, 0xe8,
	0x17, 0x25, 0x10, 0xf0, 0xa9, 0x06, 0xfd, 0x14, 0xa4, 0x2f, 0x97, 0x8b, 0x59, 0x0a, 0x29, 0x28,
	0x24, 0xa8, 0xbf, 0x34, 0x69, 0x31, 0x8f, 0x01, 0xb7, 0x80, 0xf7, 0x0e, 0x20, 0x0d, 0xd0, 0xb1,
	0x36, 0xc0, 0xf9, 0x9f, 0x01, 0x9b, 0xbc, 0xd7, 0x07, 0x7d, 0xa2, 0x88, 0x04, 0x7f, 0xc9, 0x66,
	0x48, 0x51, 0x49, 0x59, 0x9e, 0xae, 0x8b, 0x12, 0x0a, 0xc0, 0x68, 0xb3, 0xce, 0x12, 0xdb, 0x72,
	0x2d, 0x6f, 0x10, 0x72, 0x83, 0x5d, 0x35, 0xd0, 0x87, 0x84, 0xaf, 0xd8, 0x69, 0x22, 0x0a, 0xc0,
	0x8c, 0xd0, 0x7e, 0xe0, 0xf6, 0xbd, 0xf1, 0xea, 0xa9, 0xaf, 0x73, 0x34, 0x61, 0xfd, 0x77, 0x1a,
	0x0e, 0xef, 0x79, 0xfc, 0x39, 0x3b, 0x91, 0x40, 0x02, 0xed, 0xbe, 0x12, 0x3c, 0xee, 0x08, 0xbe,
	0x00, 0x89, 0x50, 0x33, 0xf8, 0x05, 0x1b, 0x99, 0x1c, 0x68, 0x0f, 0x14, 0x7d, 0xde, 0xa1, 0x9b,
	0x30, 0xe1, 0x81, 0xc9, 0x2f, 0xd9, 0xa3, 0xe6, 0xb4, 0x75, 0x11, 0x95, 0xd1, 0x16, 0xed, 0x13,
	0xd7, 0xf2, 0xc6, 0xab, 0x67, 0xff, 0xcf, 0x76, 0xa5, 0x38, 0xe1, 0x34, 0x69, 0x8f, 0xfc, 0x0d,
	0x9b, 0x4a, 0xd0, 0x57, 0xa1, 0x3d, 0x86, 0xca, 0xe3, 0xec, 0xdf, 0xb8, 0xf5, 0x95, 0x68, 0x8b,
	0x89, 0x6c, 0x4d, 0xfc, 0x35, 0x9b, 0x50, 0xb4, 0xd9, 0xdc, 0x1a, 0x83, 0x87, 0xca, 0x60, 0xd1,
	0x31, 0xf8, 0x5c, 0x53, 0x1a, 0xfd, 0x98, 0x0e, 0x03, 0x7f, 0xc1, 0x86, 0x8d, 0xf0, 0x54, 0x09,
	0x9f, 0x74, 0x37, 0xd7, 0x9a, 0x86, 0xc4, 0xcf, 0xd9, 0x24, 0x86, 0x1c, 0x29, 0xa3, 0x8a, 0x32,
	0xc8, 0xed, 0x91, 0x6b, 0x79, 0xa3, 0xf0, 0xe8, 0x5f, 0xbd, 0x53, 0x5c, 0x21, 0xc1, 0xd6, 0x44,
	0x62, 0xcd, 0x4e, 0x47, 0x05, 0xf3, 0x2f, 0x15, 0xc7, 0xec, 0x14, 0xb7, 0x26, 0x7e, 0xc1, 0xe6,
	0xb1, 0x28, 0x69, 0x5d, 0xbf, 0x4f, 0xd2, 0x2e, 0x09, 0xda, 0x63, 0xb7, 0xef, 0x0d, 0xc2, 0x59,
	0x0d, 0xd7, 0x6f, 0x98, 0x1c, 0x6a, 0x82, 0x6f, 0x3f, 0xfe, 0xdc, 0x39, 0xd6, 0xdd, 0xce, 0xb1,
	0x7e, 0xef, 0x1c, 0xeb, 0xc7, 0xde, 0xe9, 0xdd, 0xed, 0x9d, 0xde, 0xaf, 0xbd, 0xd3, 0xfb, 0xba,
	0x4c, 0x33, 0xba, 0xa9, 0xae, 0xfd, 0x18, 0xb6, 0x81, 0x4e, 0xf1, 0x0d, 0xaa, 0x3c, 0x89, 0xea,
	0xbc, 0xcd, 0x8f, 0xe0, 0xbb, 0xea, 0x2e, 0xdd, 0x16, 0x02, 0x03, 0xb9, 0xbc, 0x1e, 0xaa, 0xfa,
	0xbe, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xad, 0xee, 0x1e, 0x93, 0x34, 0x03, 0x00, 0x00,
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
	if len(m.CertVotedProposalIds) > 0 {
		dAtA2 := make([]byte, len(m.CertVotedProposalIds)*10)
		var j1 int
		for _, num := range m.CertVotedProposalIds {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGenesis(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x5a
	}
	if m.CustomParams != nil {
		{
			size, err := m.CustomParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if len(m.Constitution) > 0 {
		i -= len(m.Constitution)
		copy(dAtA[i:], m.Constitution)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Constitution)))
		i--
		dAtA[i] = 0x4a
	}
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	if m.TallyParams != nil {
		{
			size, err := m.TallyParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.VotingParams != nil {
		{
			size, err := m.VotingParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.DepositParams != nil {
		{
			size, err := m.DepositParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.StartingProposalId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingProposalId))
		i--
		dAtA[i] = 0x8
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
	if m.StartingProposalId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingProposalId))
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.DepositParams != nil {
		l = m.DepositParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.VotingParams != nil {
		l = m.VotingParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TallyParams != nil {
		l = m.TallyParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.Constitution)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.CustomParams != nil {
		l = m.CustomParams.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.CertVotedProposalIds) > 0 {
		l = 0
		for _, e := range m.CertVotedProposalIds {
			l += sovGenesis(uint64(e))
		}
		n += 1 + sovGenesis(uint64(l)) + l
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingProposalId", wireType)
			}
			m.StartingProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
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
			m.Deposits = append(m.Deposits, &v1.Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
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
			m.Votes = append(m.Votes, &v1.Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
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
			m.Proposals = append(m.Proposals, &v1.Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositParams", wireType)
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
			if m.DepositParams == nil {
				m.DepositParams = &v1.DepositParams{}
			}
			if err := m.DepositParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingParams", wireType)
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
			if m.VotingParams == nil {
				m.VotingParams = &v1.VotingParams{}
			}
			if err := m.VotingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyParams", wireType)
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
			if m.TallyParams == nil {
				m.TallyParams = &v1.TallyParams{}
			}
			if err := m.TallyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if m.Params == nil {
				m.Params = &v1.Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constitution", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Constitution = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CustomParams", wireType)
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
			if m.CustomParams == nil {
				m.CustomParams = &CustomParams{}
			}
			if err := m.CustomParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CertVotedProposalIds = append(m.CertVotedProposalIds, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGenesis
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGenesis
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGenesis
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.CertVotedProposalIds) == 0 {
					m.CertVotedProposalIds = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGenesis
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CertVotedProposalIds = append(m.CertVotedProposalIds, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CertVotedProposalIds", wireType)
			}
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
