// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/gov/v1alpha1/genesis.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/x/gov/types"
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

// GenesisState defines the gov module's genesis state.
type GenesisState struct {
	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty" yaml:"starting_proposal_id"`
	// deposits defines all the deposits present at genesis.
	Deposits []types.Deposit `protobuf:"bytes,2,rep,name=deposits,proto3" json:"deposits"`
	// votes defines all the votes present at genesis.
	Votes []types.Vote `protobuf:"bytes,3,rep,name=votes,proto3" json:"votes"`
	// proposals defines all the proposals present at genesis.
	Proposals []types.Proposal `protobuf:"bytes,4,rep,name=proposals,proto3" json:"proposals"`
	// params defines all the parameters of related to deposit.
	DepositParams types.DepositParams `protobuf:"bytes,5,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params" yaml:"deposit_params"`
	// params defines all the parameters of related to voting.
	VotingParams types.VotingParams `protobuf:"bytes,6,opt,name=voting_params,json=votingParams,proto3" json:"voting_params" yaml:"voting_params"`
	// params defines all the parameters of related to tally.
	TallyParams types.TallyParams `protobuf:"bytes,7,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params" yaml:"tally_params"`
	// params defines all the parameters of related to custom.
	CustomParams CustomParams `protobuf:"bytes,8,opt,name=custom_params,json=customParams,proto3" json:"custom_params" yaml:"custom_params"`
	// proposals that require and have passed cert votes.
	CertVotedProposalIds []uint64 `protobuf:"varint,9,rep,packed,name=cert_voted_proposal_ids,json=certVotedProposalIds,proto3" json:"cert_voted_proposal_ids,omitempty" yaml:"cert_voted_proposal_ids"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d2eda471d71961d, []int{0}
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

func (m *GenesisState) GetDeposits() []types.Deposit {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetVotes() []types.Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisState) GetProposals() []types.Proposal {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisState) GetDepositParams() types.DepositParams {
	if m != nil {
		return m.DepositParams
	}
	return types.DepositParams{}
}

func (m *GenesisState) GetVotingParams() types.VotingParams {
	if m != nil {
		return m.VotingParams
	}
	return types.VotingParams{}
}

func (m *GenesisState) GetTallyParams() types.TallyParams {
	if m != nil {
		return m.TallyParams
	}
	return types.TallyParams{}
}

func (m *GenesisState) GetCustomParams() CustomParams {
	if m != nil {
		return m.CustomParams
	}
	return CustomParams{}
}

func (m *GenesisState) GetCertVotedProposalIds() []uint64 {
	if m != nil {
		return m.CertVotedProposalIds
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "shentu.gov.v1alpha1.GenesisState")
}

func init() { proto.RegisterFile("shentu/gov/v1alpha1/genesis.proto", fileDescriptor_1d2eda471d71961d) }

var fileDescriptor_1d2eda471d71961d = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x63, 0x92, 0x86, 0x76, 0x93, 0x70, 0xd8, 0x06, 0x61, 0x35, 0xa9, 0x9d, 0xfa, 0x94,
	0x93, 0xad, 0x16, 0x4e, 0x48, 0x48, 0xc8, 0x20, 0x01, 0xb7, 0x62, 0x10, 0x12, 0x5c, 0xac, 0x8d,
	0xbd, 0x38, 0x96, 0x6c, 0x8f, 0xe5, 0xdd, 0x58, 0xe4, 0x0d, 0x38, 0xf2, 0x58, 0x3d, 0xf6, 0xc8,
	0x29, 0x42, 0xc9, 0x1b, 0xe4, 0x09, 0x90, 0x77, 0xd7, 0xc4, 0x91, 0x4c, 0x6f, 0xd1, 0xce, 0xff,
	0xff, 0xdf, 0xcc, 0x64, 0x8c, 0xae, 0xd8, 0x92, 0x66, 0x7c, 0xe5, 0x44, 0x50, 0x3a, 0xe5, 0x35,
	0x49, 0xf2, 0x25, 0xb9, 0x76, 0x22, 0x9a, 0x51, 0x16, 0x33, 0x3b, 0x2f, 0x80, 0x03, 0x3e, 0x97,
	0x12, 0x3b, 0x82, 0xd2, 0xae, 0x25, 0x17, 0xe3, 0x08, 0x22, 0x10, 0x75, 0xa7, 0xfa, 0x25, 0xa5,
	0x17, 0xd3, 0x00, 0x58, 0x0a, 0x4c, 0xa5, 0x2d, 0x28, 0xaf, 0xc2, 0xa0, 0x54, 0xd5, 0xcb, 0x56,
	0x56, 0x5d, 0xb6, 0x7e, 0xf6, 0xd1, 0xf0, 0x9d, 0x24, 0x7f, 0xe2, 0x84, 0x53, 0xfc, 0x11, 0x8d,
	0x19, 0x27, 0x05, 0x8f, 0xb3, 0xc8, 0xcf, 0x0b, 0xc8, 0x81, 0x91, 0xc4, 0x8f, 0x43, 0x5d, 0x9b,
	0x69, 0xf3, 0x9e, 0x6b, 0xee, 0x37, 0xe6, 0x64, 0x4d, 0xd2, 0xe4, 0xa5, 0xd5, 0xa6, 0xb2, 0x3c,
	0x5c, 0x3f, 0xdf, 0xaa, 0xd7, 0x0f, 0x21, 0x7e, 0x85, 0x4e, 0x43, 0x9a, 0x03, 0x8b, 0x39, 0xd3,
	0x1f, 0xcd, 0xba, 0xf3, 0xc1, 0xcd, 0xc4, 0x96, 0x3d, 0xab, 0xf1, 0x44, 0xcf, 0xf6, 0x5b, 0xa9,
	0x71, 0x7b, 0x77, 0x1b, 0xb3, 0xe3, 0xfd, 0xb3, 0xe0, 0x17, 0xe8, 0xa4, 0x04, 0x4e, 0x99, 0xde,
	0x15, 0x5e, 0xbd, 0xcd, 0xfb, 0x05, 0x38, 0x55, 0x46, 0x29, 0xc6, 0xaf, 0xd1, 0x59, 0xdd, 0x18,
	0xd3, 0x7b, 0xc2, 0x39, 0x6d, 0x73, 0xd6, 0x7d, 0x2a, 0xf7, 0xc1, 0x84, 0x23, 0xf4, 0x44, 0xf5,
	0xe0, 0xe7, 0xa4, 0x20, 0x29, 0xd3, 0x4f, 0x66, 0xda, 0x7c, 0x70, 0x73, 0xf5, 0x40, 0xf3, 0xb7,
	0x42, 0xe8, 0x5e, 0x56, 0x59, 0xfb, 0x8d, 0xf9, 0x54, 0xae, 0xea, 0x38, 0xc6, 0xf2, 0x46, 0x61,
	0x53, 0x8d, 0x03, 0x34, 0x2a, 0x41, 0xae, 0x52, 0x72, 0xfa, 0x82, 0x33, 0xfb, 0xcf, 0xa0, 0xd5,
	0x72, 0x25, 0x66, 0xaa, 0x30, 0x63, 0x89, 0x39, 0x0a, 0xb1, 0xbc, 0x61, 0xd9, 0xd0, 0x62, 0x1f,
	0x0d, 0x39, 0x49, 0x92, 0x75, 0xcd, 0x78, 0x2c, 0x18, 0x66, 0x1b, 0xe3, 0x73, 0xa5, 0x53, 0x88,
	0x89, 0x42, 0x9c, 0x4b, 0x44, 0x33, 0xc2, 0xf2, 0x06, 0xfc, 0xa0, 0xc4, 0x21, 0x1a, 0x05, 0x2b,
	0xc6, 0x21, 0xad, 0x09, 0xa7, 0x6a, 0x5b, 0x2d, 0x97, 0x6c, 0xbf, 0x11, 0xca, 0xf6, 0x31, 0x8e,
	0x52, 0x2c, 0x6f, 0x18, 0x34, 0xb4, 0xf8, 0x2b, 0x7a, 0x16, 0xd0, 0x82, 0xfb, 0xd5, 0x9f, 0x1c,
	0x36, 0x4f, 0x8f, 0xe9, 0x67, 0xb3, 0xee, 0xbc, 0xe7, 0x5a, 0xfb, 0x8d, 0x69, 0xa8, 0xa0, 0x76,
	0xa1, 0xe5, 0x8d, 0xab, 0x4a, 0x75, 0x32, 0xe1, 0xe1, 0x4a, 0x99, 0xfb, 0xfe, 0x6e, 0x6b, 0x68,
	0xf7, 0x5b, 0x43, 0xfb, 0xb3, 0x35, 0xb4, 0x5f, 0x3b, 0xa3, 0x73, 0xbf, 0x33, 0x3a, 0xbf, 0x77,
	0x46, 0xe7, 0x9b, 0x1d, 0xc5, 0x7c, 0xb9, 0x5a, 0xd8, 0x01, 0xa4, 0x8e, 0x9c, 0xe6, 0x3b, 0xac,
	0xb2, 0x90, 0xf0, 0x18, 0x32, 0xf5, 0xe0, 0xfc, 0x10, 0x5f, 0x18, 0x5f, 0xe7, 0x94, 0x2d, 0xfa,
	0xe2, 0xdb, 0x7a, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xe4, 0x06, 0xe3, 0xe8, 0x03, 0x00,
	0x00,
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
		dAtA[i] = 0x4a
	}
	{
		size, err := m.CustomParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
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
	l = m.DepositParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.VotingParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TallyParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.CustomParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
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
			m.Deposits = append(m.Deposits, types.Deposit{})
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
			m.Votes = append(m.Votes, types.Vote{})
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
			m.Proposals = append(m.Proposals, types.Proposal{})
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
			if err := m.TallyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
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
			if err := m.CustomParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
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
