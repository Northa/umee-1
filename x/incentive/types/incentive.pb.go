// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: umee/incentive/v1/incentive.proto

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

// Params defines the parameters for the incentive module.
type Params struct {
	// max_unbondings defines the maximum amount of concurrent unbondings an address can have.
	MaxUnbondings uint32 `protobuf:"varint,1,opt,name=max_unbondings,json=maxUnbondings,proto3" json:"max_unbondings,omitempty"`
	// unbonding_duration_long defines the unbonding duration (in blocks) of the long tier.
	UnbondingDurationLong uint64 `protobuf:"varint,2,opt,name=unbonding_duration_long,json=unbondingDurationLong,proto3" json:"unbonding_duration_long,omitempty"`
	// unbonding_duration_middle defines the unbonding duration (in blocks) of the middle tier.
	UnbondingDurationMiddle uint64 `protobuf:"varint,3,opt,name=unbonding_duration_middle,json=unbondingDurationMiddle,proto3" json:"unbonding_duration_middle,omitempty"`
	// unbonding_duration_short defines the unbonding duration (in blocks) of the short tier.
	UnbondingDurationShort uint64 `protobuf:"varint,4,opt,name=unbonding_duration_short,json=unbondingDurationShort,proto3" json:"unbonding_duration_short,omitempty"`
	// The tier_weight_short defines how the proportion of rewards assets locked
	// in the short unbonding duration receive compared to what the same amount
	// would receive on the long tier.
	TierWeightShort github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=tier_weight_short,json=tierWeightShort,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tier_weight_short"`
	// The tier_weight_middle defines how the proportion of rewards assets locked
	// in the middle unbonding duration receive compared to what the same amount
	// would receive on the long tier.
	TierWeightMiddle github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=tier_weight_middle,json=tierWeightMiddle,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tier_weight_middle"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c99c623956e199b, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxUnbondings() uint32 {
	if m != nil {
		return m.MaxUnbondings
	}
	return 0
}

func (m *Params) GetUnbondingDurationLong() uint64 {
	if m != nil {
		return m.UnbondingDurationLong
	}
	return 0
}

func (m *Params) GetUnbondingDurationMiddle() uint64 {
	if m != nil {
		return m.UnbondingDurationMiddle
	}
	return 0
}

func (m *Params) GetUnbondingDurationShort() uint64 {
	if m != nil {
		return m.UnbondingDurationShort
	}
	return 0
}

// IncentiveProgram defines a liquidity mining incentive program on a single
// locked uToken denom that will run for a set amount of time.
type IncentiveProgram struct {
	// ID uniquely identifies the incentive program after it has been created.
	// It is zero when the program is being proposed by governance.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// start_height is the block height at which the incentives begin.
	StartHeight uint64 `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	// duration is the length of the incentive program in blocks.
	Duration uint64 `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	// lock_denom defines the denomination of the incentivized collateral uToken.
	LockDenom string `protobuf:"bytes,4,opt,name=lock_denom,json=lockDenom,proto3" json:"lock_denom,omitempty"`
	// total_rewards are total amount of rewards which can be distributed to
	// suppliers by this program. This is set to its final value when the program
	// is proposed by governance.
	TotalRewards types.Coin `protobuf:"bytes,5,opt,name=total_rewards,json=totalRewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"total_rewards"`
	// funded_rewards are total amount of rewards which have been funded by a
	// sponsor to this program. This is zero until the program is both passed
	// by governance and funded.
	FundedRewards types.Coin `protobuf:"bytes,6,opt,name=funded_rewards,json=fundedRewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"funded_rewards"`
	// remaining_rewards are total amount of this program's funded rewards
	// which have not yet been allocated to suppliers. This is zero until the
	// program is both passed by governance and funded, then begins decreasing
	// to zero as the program runs to completion.
	RemainingRewards types.Coin `protobuf:"bytes,7,opt,name=remaining_rewards,json=remainingRewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"remaining_rewards"`
}

func (m *IncentiveProgram) Reset()         { *m = IncentiveProgram{} }
func (m *IncentiveProgram) String() string { return proto.CompactTextString(m) }
func (*IncentiveProgram) ProtoMessage()    {}
func (*IncentiveProgram) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c99c623956e199b, []int{1}
}
func (m *IncentiveProgram) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IncentiveProgram) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IncentiveProgram.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IncentiveProgram) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncentiveProgram.Merge(m, src)
}
func (m *IncentiveProgram) XXX_Size() int {
	return m.Size()
}
func (m *IncentiveProgram) XXX_DiscardUnknown() {
	xxx_messageInfo_IncentiveProgram.DiscardUnknown(m)
}

var xxx_messageInfo_IncentiveProgram proto.InternalMessageInfo

func (m *IncentiveProgram) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *IncentiveProgram) GetStartHeight() uint64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *IncentiveProgram) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *IncentiveProgram) GetLockDenom() string {
	if m != nil {
		return m.LockDenom
	}
	return ""
}

func (m *IncentiveProgram) GetTotalRewards() types.Coin {
	if m != nil {
		return m.TotalRewards
	}
	return types.Coin{}
}

func (m *IncentiveProgram) GetFundedRewards() types.Coin {
	if m != nil {
		return m.FundedRewards
	}
	return types.Coin{}
}

func (m *IncentiveProgram) GetRemainingRewards() types.Coin {
	if m != nil {
		return m.RemainingRewards
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*Params)(nil), "umeenetwork.umee.incentive.v1.Params")
	proto.RegisterType((*IncentiveProgram)(nil), "umeenetwork.umee.incentive.v1.IncentiveProgram")
}

func init() { proto.RegisterFile("umee/incentive/v1/incentive.proto", fileDescriptor_8c99c623956e199b) }

var fileDescriptor_8c99c623956e199b = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4f, 0x6b, 0x13, 0x41,
	0x1c, 0xcd, 0x26, 0x31, 0x9a, 0x69, 0x13, 0xd3, 0x41, 0x6d, 0x1a, 0xe8, 0x26, 0x2d, 0x28, 0xb9,
	0x74, 0xd7, 0x58, 0x10, 0xe9, 0x31, 0xe6, 0xa0, 0xa8, 0x50, 0x56, 0x44, 0x28, 0xc2, 0x32, 0xd9,
	0x19, 0x37, 0x43, 0xb2, 0x33, 0x61, 0x66, 0xf2, 0xc7, 0x6f, 0xe1, 0xcd, 0xab, 0x67, 0x3f, 0x49,
	0x8f, 0x3d, 0x8a, 0x87, 0x2a, 0x09, 0x88, 0x1f, 0x43, 0x66, 0x76, 0xb2, 0x09, 0xb4, 0x07, 0x91,
	0x9e, 0x76, 0xe6, 0xfd, 0xe6, 0xbd, 0xc7, 0xef, 0xdf, 0x82, 0x83, 0x49, 0x42, 0x88, 0x4f, 0x59,
	0x44, 0x98, 0xa2, 0x53, 0xe2, 0x4f, 0x3b, 0xeb, 0x8b, 0x37, 0x16, 0x5c, 0x71, 0xb8, 0xaf, 0x9f,
	0x30, 0xa2, 0x66, 0x5c, 0x0c, 0x3d, 0x7d, 0xf6, 0xd6, 0x2f, 0xa6, 0x9d, 0x86, 0x1b, 0x71, 0x99,
	0x70, 0xe9, 0xf7, 0x91, 0xd4, 0xf4, 0x3e, 0x51, 0xa8, 0xe3, 0x47, 0x9c, 0xb2, 0x94, 0xde, 0xb8,
	0x17, 0xf3, 0x98, 0x9b, 0xa3, 0xaf, 0x4f, 0x29, 0x7a, 0xf8, 0xa5, 0x00, 0x4a, 0xa7, 0x48, 0xa0,
	0x44, 0xc2, 0x87, 0xa0, 0x9a, 0xa0, 0x79, 0x38, 0x61, 0x7d, 0xce, 0x30, 0x65, 0xb1, 0xac, 0x3b,
	0x2d, 0xa7, 0x5d, 0x09, 0x2a, 0x09, 0x9a, 0xbf, 0xcb, 0x40, 0xf8, 0x14, 0xec, 0x66, 0x4f, 0x42,
	0x3c, 0x11, 0x48, 0x51, 0xce, 0xc2, 0x11, 0x67, 0x71, 0x3d, 0xdf, 0x72, 0xda, 0xc5, 0xe0, 0x7e,
	0x16, 0xee, 0xd9, 0xe8, 0x6b, 0xce, 0x62, 0x78, 0x02, 0xf6, 0xae, 0xe1, 0x25, 0x14, 0xe3, 0x11,
	0xa9, 0x17, 0x0c, 0x73, 0xf7, 0x0a, 0xf3, 0x8d, 0x09, 0xc3, 0x67, 0xa0, 0x7e, 0x0d, 0x57, 0x0e,
	0xb8, 0x50, 0xf5, 0xa2, 0xa1, 0x3e, 0xb8, 0x42, 0x7d, 0xab, 0xa3, 0xf0, 0x0c, 0xec, 0x28, 0x4a,
	0x44, 0x38, 0x23, 0x34, 0x1e, 0x28, 0x4b, 0xb9, 0xd5, 0x72, 0xda, 0xe5, 0xae, 0x77, 0x7e, 0xd9,
	0xcc, 0xfd, 0xb8, 0x6c, 0x3e, 0x8a, 0xa9, 0x1a, 0x4c, 0xfa, 0x5e, 0xc4, 0x13, 0xdf, 0xd6, 0x30,
	0xfd, 0x1c, 0x49, 0x3c, 0xf4, 0xd5, 0xa7, 0x31, 0x91, 0x5e, 0x8f, 0x44, 0xc1, 0x5d, 0x2d, 0xf4,
	0xde, 0xe8, 0xa4, 0xda, 0x1f, 0x00, 0xdc, 0xd4, 0xb6, 0xa9, 0x94, 0xfe, 0x4b, 0xbc, 0xb6, 0x16,
	0x4f, 0x73, 0x3e, 0xfc, 0x5d, 0x00, 0xb5, 0x97, 0xab, 0x06, 0x9f, 0x0a, 0x1e, 0x0b, 0x94, 0xc0,
	0x2a, 0xc8, 0x53, 0x6c, 0xfb, 0x92, 0xa7, 0x18, 0x1e, 0x80, 0x6d, 0xa9, 0x90, 0x50, 0xe1, 0xc0,
	0x50, 0x6d, 0x07, 0xb6, 0x0c, 0xf6, 0xc2, 0x40, 0xb0, 0x01, 0xee, 0xac, 0x2a, 0x66, 0xcb, 0x9c,
	0xdd, 0xe1, 0x3e, 0x00, 0x23, 0x1e, 0x0d, 0x43, 0x4c, 0x18, 0x4f, 0x4c, 0x25, 0xcb, 0x41, 0x59,
	0x23, 0x3d, 0x0d, 0xc0, 0x31, 0xa8, 0x28, 0xae, 0xd0, 0x28, 0x14, 0x64, 0x86, 0x04, 0x96, 0xa6,
	0x70, 0x5b, 0x4f, 0xf6, 0xbc, 0x34, 0x05, 0x4f, 0x8f, 0x9a, 0x67, 0x47, 0xcd, 0x7b, 0xce, 0x29,
	0xeb, 0x3e, 0xd6, 0x69, 0x7f, 0xfb, 0xd9, 0x6c, 0xff, 0x43, 0xda, 0x9a, 0x20, 0x83, 0x6d, 0xe3,
	0x10, 0xa4, 0x06, 0x50, 0x80, 0xea, 0xc7, 0x09, 0xc3, 0x04, 0x67, 0x96, 0xa5, 0x9b, 0xb7, 0xac,
	0xa4, 0x16, 0x2b, 0xcf, 0x39, 0xd8, 0x11, 0x24, 0x41, 0x94, 0xe9, 0xe1, 0x5a, 0xd9, 0xde, 0xbe,
	0x79, 0xdb, 0x5a, 0xe6, 0x62, 0x9d, 0x4f, 0x8a, 0x7f, 0xbe, 0x36, 0x9d, 0xee, 0xab, 0xf3, 0x85,
	0xeb, 0x5c, 0x2c, 0x5c, 0xe7, 0xd7, 0xc2, 0x75, 0x3e, 0x2f, 0xdd, 0xdc, 0xc5, 0xd2, 0xcd, 0x7d,
	0x5f, 0xba, 0xb9, 0xb3, 0xce, 0x86, 0xb6, 0x5e, 0xf8, 0x23, 0xbb, 0xfd, 0xe6, 0xe2, 0x4f, 0x8f,
	0xfd, 0xf9, 0xc6, 0x1f, 0xc3, 0x58, 0xf5, 0x4b, 0x66, 0xad, 0x8f, 0xff, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x0e, 0xb8, 0xc0, 0x11, 0x50, 0x04, 0x00, 0x00,
}

func (this *IncentiveProgram) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IncentiveProgram)
	if !ok {
		that2, ok := that.(IncentiveProgram)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.StartHeight != that1.StartHeight {
		return false
	}
	if this.Duration != that1.Duration {
		return false
	}
	if this.LockDenom != that1.LockDenom {
		return false
	}
	if !this.TotalRewards.Equal(&that1.TotalRewards) {
		return false
	}
	if !this.FundedRewards.Equal(&that1.FundedRewards) {
		return false
	}
	if !this.RemainingRewards.Equal(&that1.RemainingRewards) {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TierWeightMiddle.Size()
		i -= size
		if _, err := m.TierWeightMiddle.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.TierWeightShort.Size()
		i -= size
		if _, err := m.TierWeightShort.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.UnbondingDurationShort != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.UnbondingDurationShort))
		i--
		dAtA[i] = 0x20
	}
	if m.UnbondingDurationMiddle != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.UnbondingDurationMiddle))
		i--
		dAtA[i] = 0x18
	}
	if m.UnbondingDurationLong != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.UnbondingDurationLong))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxUnbondings != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.MaxUnbondings))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *IncentiveProgram) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IncentiveProgram) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IncentiveProgram) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.RemainingRewards.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.FundedRewards.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.TotalRewards.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIncentive(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.LockDenom) > 0 {
		i -= len(m.LockDenom)
		copy(dAtA[i:], m.LockDenom)
		i = encodeVarintIncentive(dAtA, i, uint64(len(m.LockDenom)))
		i--
		dAtA[i] = 0x22
	}
	if m.Duration != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x18
	}
	if m.StartHeight != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintIncentive(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIncentive(dAtA []byte, offset int, v uint64) int {
	offset -= sovIncentive(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxUnbondings != 0 {
		n += 1 + sovIncentive(uint64(m.MaxUnbondings))
	}
	if m.UnbondingDurationLong != 0 {
		n += 1 + sovIncentive(uint64(m.UnbondingDurationLong))
	}
	if m.UnbondingDurationMiddle != 0 {
		n += 1 + sovIncentive(uint64(m.UnbondingDurationMiddle))
	}
	if m.UnbondingDurationShort != 0 {
		n += 1 + sovIncentive(uint64(m.UnbondingDurationShort))
	}
	l = m.TierWeightShort.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.TierWeightMiddle.Size()
	n += 1 + l + sovIncentive(uint64(l))
	return n
}

func (m *IncentiveProgram) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovIncentive(uint64(m.Id))
	}
	if m.StartHeight != 0 {
		n += 1 + sovIncentive(uint64(m.StartHeight))
	}
	if m.Duration != 0 {
		n += 1 + sovIncentive(uint64(m.Duration))
	}
	l = len(m.LockDenom)
	if l > 0 {
		n += 1 + l + sovIncentive(uint64(l))
	}
	l = m.TotalRewards.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.FundedRewards.Size()
	n += 1 + l + sovIncentive(uint64(l))
	l = m.RemainingRewards.Size()
	n += 1 + l + sovIncentive(uint64(l))
	return n
}

func sovIncentive(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIncentive(x uint64) (n int) {
	return sovIncentive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxUnbondings", wireType)
			}
			m.MaxUnbondings = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxUnbondings |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingDurationLong", wireType)
			}
			m.UnbondingDurationLong = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingDurationLong |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingDurationMiddle", wireType)
			}
			m.UnbondingDurationMiddle = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingDurationMiddle |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnbondingDurationShort", wireType)
			}
			m.UnbondingDurationShort = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UnbondingDurationShort |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TierWeightShort", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TierWeightShort.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TierWeightMiddle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TierWeightMiddle.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func (m *IncentiveProgram) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIncentive
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
			return fmt.Errorf("proto: IncentiveProgram: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IncentiveProgram: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LockDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LockDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FundedRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FundedRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RemainingRewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIncentive
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
				return ErrInvalidLengthIncentive
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIncentive
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RemainingRewards.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIncentive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIncentive
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
func skipIncentive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
					return 0, ErrIntOverflowIncentive
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
				return 0, ErrInvalidLengthIncentive
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIncentive
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIncentive
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIncentive        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIncentive          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIncentive = fmt.Errorf("proto: unexpected end of group")
)
