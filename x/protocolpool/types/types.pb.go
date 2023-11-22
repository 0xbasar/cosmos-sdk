// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/protocolpool/v1/types.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
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

// Budget defines the fields of a budget proposal.
type Budget struct {
	// recipient_address is the address of the recipient who can claim the budget.
	RecipientAddress string `protobuf:"bytes,1,opt,name=recipient_address,json=recipientAddress,proto3" json:"recipient_address,omitempty"`
	// total_budget is the total amount allocated for the budget.
	TotalBudget *types.Coin `protobuf:"bytes,2,opt,name=total_budget,json=totalBudget,proto3" json:"total_budget,omitempty"`
	// claimed_amount is the total amount claimed from the total budget amount requested.
	ClaimedAmount *types.Coin `protobuf:"bytes,3,opt,name=claimed_amount,json=claimedAmount,proto3" json:"claimed_amount,omitempty"`
	// start_time is the time when the budget becomes claimable.
	StartTime *time.Time `protobuf:"bytes,4,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time,omitempty"`
	// next_claim_from is the time when the budget was last successfully claimed or distributed.
	// It is used to track the next starting claim time for fund distribution. If set, it cannot be less than start_time.
	NextClaimFrom *time.Time `protobuf:"bytes,5,opt,name=next_claim_from,json=nextClaimFrom,proto3,stdtime" json:"next_claim_from,omitempty"`
	// tranches is the number of times the total budget amount is to be distributed.
	Tranches uint64 `protobuf:"varint,6,opt,name=tranches,proto3" json:"tranches,omitempty"`
	// tranches_left is the number of tranches left for the amount to be distributed.
	TranchesLeft uint64 `protobuf:"varint,7,opt,name=tranches_left,json=tranchesLeft,proto3" json:"tranches_left,omitempty"`
	// Period is the time interval(number of seconds) at which funds distribution should be performed.
	// For example, if a period is set to 3600, it represents an action that
	// should occur every hour (3600 seconds).
	Period *time.Duration `protobuf:"bytes,8,opt,name=period,proto3,stdduration" json:"period,omitempty"`
}

func (m *Budget) Reset()         { *m = Budget{} }
func (m *Budget) String() string { return proto.CompactTextString(m) }
func (*Budget) ProtoMessage()    {}
func (*Budget) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b7d0ea246d7f44, []int{0}
}
func (m *Budget) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Budget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Budget.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Budget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Budget.Merge(m, src)
}
func (m *Budget) XXX_Size() int {
	return m.Size()
}
func (m *Budget) XXX_DiscardUnknown() {
	xxx_messageInfo_Budget.DiscardUnknown(m)
}

var xxx_messageInfo_Budget proto.InternalMessageInfo

func (m *Budget) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *Budget) GetTotalBudget() *types.Coin {
	if m != nil {
		return m.TotalBudget
	}
	return nil
}

func (m *Budget) GetClaimedAmount() *types.Coin {
	if m != nil {
		return m.ClaimedAmount
	}
	return nil
}

func (m *Budget) GetStartTime() *time.Time {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Budget) GetNextClaimFrom() *time.Time {
	if m != nil {
		return m.NextClaimFrom
	}
	return nil
}

func (m *Budget) GetTranches() uint64 {
	if m != nil {
		return m.Tranches
	}
	return 0
}

func (m *Budget) GetTranchesLeft() uint64 {
	if m != nil {
		return m.TranchesLeft
	}
	return 0
}

func (m *Budget) GetPeriod() *time.Duration {
	if m != nil {
		return m.Period
	}
	return nil
}

// ContinuousFund defines the fields of continuous fund proposal.
type ContinuousFund struct {
	// Recipient address of the account receiving funds.
	Recipient string `protobuf:"bytes,1,opt,name=recipient,proto3" json:"recipient,omitempty"`
	// Percentage is the percentage of funds to be allocated from Community pool share on block by block,
	// till the `cap` is reached or expired.
	Percentage cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=percentage,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"percentage"`
	// Cap is the capital amount, which when its met funds are no longer distributed.
	Cap *types.Coin `protobuf:"bytes,3,opt,name=cap,proto3" json:"cap,omitempty"`
	// Optional, if expiry is set, removes the state object when expired.
	Expiry *time.Time `protobuf:"bytes,4,opt,name=expiry,proto3,stdtime" json:"expiry,omitempty"`
}

func (m *ContinuousFund) Reset()         { *m = ContinuousFund{} }
func (m *ContinuousFund) String() string { return proto.CompactTextString(m) }
func (*ContinuousFund) ProtoMessage()    {}
func (*ContinuousFund) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1b7d0ea246d7f44, []int{1}
}
func (m *ContinuousFund) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContinuousFund) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContinuousFund.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContinuousFund) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousFund.Merge(m, src)
}
func (m *ContinuousFund) XXX_Size() int {
	return m.Size()
}
func (m *ContinuousFund) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousFund.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousFund proto.InternalMessageInfo

func (m *ContinuousFund) GetRecipient() string {
	if m != nil {
		return m.Recipient
	}
	return ""
}

func (m *ContinuousFund) GetCap() *types.Coin {
	if m != nil {
		return m.Cap
	}
	return nil
}

func (m *ContinuousFund) GetExpiry() *time.Time {
	if m != nil {
		return m.Expiry
	}
	return nil
}

func init() {
	proto.RegisterType((*Budget)(nil), "cosmos.protocolpool.v1.Budget")
	proto.RegisterType((*ContinuousFund)(nil), "cosmos.protocolpool.v1.ContinuousFund")
}

func init() {
	proto.RegisterFile("cosmos/protocolpool/v1/types.proto", fileDescriptor_c1b7d0ea246d7f44)
}

var fileDescriptor_c1b7d0ea246d7f44 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0xdb, 0x7c, 0xf9, 0x9a, 0xe9, 0x0f, 0x30, 0xaa, 0x90, 0x1b, 0x24, 0x27, 0x84, 0x4d,
	0x24, 0x54, 0x5b, 0x01, 0x09, 0x90, 0x40, 0x82, 0x3a, 0xa1, 0x62, 0xd1, 0x0d, 0x81, 0x15, 0x1b,
	0x6b, 0x62, 0xdf, 0xb8, 0x23, 0x6c, 0x5f, 0x6b, 0x3c, 0x8e, 0x92, 0x2d, 0x4f, 0xd0, 0x25, 0x0f,
	0xd2, 0x87, 0xe8, 0xb2, 0xaa, 0x58, 0x20, 0x16, 0x05, 0x25, 0x2f, 0x82, 0x3c, 0x9e, 0x84, 0xb4,
	0x9b, 0x74, 0x37, 0x3e, 0xf7, 0x9c, 0xe3, 0x33, 0xf7, 0x68, 0x48, 0xdb, 0xc7, 0x2c, 0xc6, 0xcc,
	0x49, 0x05, 0x4a, 0xf4, 0x31, 0x4a, 0x11, 0x23, 0x67, 0xdc, 0x75, 0xe4, 0x34, 0x85, 0xcc, 0x56,
	0x28, 0x7d, 0x58, 0x72, 0xec, 0x55, 0x8e, 0x3d, 0xee, 0x36, 0xf6, 0x43, 0x0c, 0x51, 0x81, 0x4e,
	0x71, 0x2a, 0xe7, 0x8d, 0x83, 0x92, 0xed, 0x95, 0x83, 0x55, 0x69, 0xc3, 0xd2, 0x3f, 0x1b, 0xb2,
	0x0c, 0x9c, 0x71, 0x77, 0x08, 0x92, 0x75, 0x1d, 0x1f, 0x79, 0xa2, 0xe7, 0xcd, 0x10, 0x31, 0x8c,
	0xa0, 0x0c, 0x33, 0xcc, 0x47, 0x8e, 0xe4, 0x31, 0x64, 0x92, 0xc5, 0xe9, 0xc2, 0xe0, 0x36, 0x21,
	0xc8, 0x05, 0x93, 0x1c, 0xb5, 0x41, 0xfb, 0xc7, 0x26, 0xa9, 0xb9, 0x79, 0x10, 0x82, 0xa4, 0xef,
	0xc9, 0x03, 0x01, 0x3e, 0x4f, 0x39, 0x24, 0xd2, 0x63, 0x41, 0x20, 0x20, 0xcb, 0x4c, 0xa3, 0x65,
	0x74, 0xea, 0xae, 0x79, 0x75, 0x7e, 0xb8, 0xaf, 0x83, 0x1d, 0x95, 0x93, 0x4f, 0x52, 0xf0, 0x24,
	0x1c, 0xdc, 0x5f, 0x4a, 0x34, 0x4e, 0xdf, 0x90, 0x1d, 0x89, 0x92, 0x45, 0xde, 0x50, 0xd9, 0x9a,
	0x1b, 0x2d, 0xa3, 0xb3, 0xfd, 0xec, 0xc0, 0xd6, 0xf2, 0xe2, 0x26, 0xb6, 0xbe, 0x89, 0xdd, 0x43,
	0x9e, 0x0c, 0xb6, 0x15, 0x5d, 0x87, 0x78, 0x47, 0xf6, 0xfc, 0x88, 0xf1, 0x18, 0x02, 0x8f, 0xc5,
	0x98, 0x27, 0xd2, 0xdc, 0x5c, 0xa7, 0xdf, 0xd5, 0x82, 0x23, 0xc5, 0xa7, 0x6f, 0x09, 0xc9, 0x24,
	0x13, 0xd2, 0x2b, 0x56, 0x61, 0x56, 0x95, 0xba, 0x61, 0x97, 0x6b, 0xb0, 0x17, 0x6b, 0xb0, 0x3f,
	0x2f, 0xf6, 0xe4, 0x56, 0xcf, 0x7e, 0x37, 0x8d, 0x41, 0x5d, 0x69, 0x0a, 0x94, 0x7e, 0x20, 0xf7,
	0x12, 0x98, 0x48, 0x4f, 0xd9, 0x7a, 0x23, 0x81, 0xb1, 0xf9, 0xdf, 0x1d, 0x5d, 0x76, 0x0b, 0x61,
	0xaf, 0xd0, 0x1d, 0x0b, 0x8c, 0x69, 0x83, 0x6c, 0x49, 0xc1, 0x12, 0xff, 0x14, 0x32, 0xb3, 0xd6,
	0x32, 0x3a, 0xd5, 0xc1, 0xf2, 0x9b, 0x3e, 0x21, 0xbb, 0x8b, 0xb3, 0x17, 0xc1, 0x48, 0x9a, 0xff,
	0x2b, 0xc2, 0xce, 0x02, 0x3c, 0x81, 0x91, 0xa4, 0x2f, 0x49, 0x2d, 0x05, 0xc1, 0x31, 0x30, 0xb7,
	0xf4, 0x16, 0x6e, 0x27, 0xe8, 0xeb, 0x3a, 0xdd, 0xea, 0xf7, 0x22, 0x80, 0xa6, 0xb7, 0xbf, 0x6d,
	0x90, 0xbd, 0x1e, 0x26, 0x92, 0x27, 0x39, 0xe6, 0xd9, 0x71, 0x9e, 0x04, 0xf4, 0x05, 0xa9, 0x2f,
	0xbb, 0x5a, 0x5b, 0xeb, 0x3f, 0x2a, 0xfd, 0x48, 0x48, 0x0a, 0xc2, 0x87, 0x44, 0xb2, 0x10, 0x54,
	0x9b, 0x75, 0xb7, 0x7b, 0x71, 0xdd, 0xac, 0xfc, 0xba, 0x6e, 0x3e, 0x2a, 0xc5, 0x59, 0xf0, 0xd5,
	0xe6, 0xe8, 0xc4, 0x4c, 0x9e, 0xda, 0x27, 0x10, 0x32, 0x7f, 0xda, 0x07, 0xff, 0xea, 0xfc, 0x90,
	0x68, 0xef, 0x3e, 0xf8, 0x83, 0x15, 0x13, 0xfa, 0x94, 0x6c, 0xfa, 0x2c, 0x5d, 0xdf, 0x6c, 0xc1,
	0xa2, 0xaf, 0x48, 0x0d, 0x26, 0x29, 0x17, 0xd3, 0x3b, 0x77, 0xa9, 0xf9, 0xee, 0xeb, 0x8b, 0x99,
	0x65, 0x5c, 0xce, 0x2c, 0xe3, 0xcf, 0xcc, 0x32, 0xce, 0xe6, 0x56, 0xe5, 0x72, 0x6e, 0x55, 0x7e,
	0xce, 0xad, 0xca, 0x97, 0xc7, 0x37, 0x72, 0x4f, 0x6e, 0xbe, 0x65, 0xf5, 0x90, 0x87, 0x35, 0x85,
	0x3d, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x63, 0x1f, 0x9d, 0x0b, 0xef, 0x03, 0x00, 0x00,
}

func (m *Budget) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Budget) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Budget) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Period != nil {
		n1, err1 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(*m.Period, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(*m.Period):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintTypes(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0x42
	}
	if m.TranchesLeft != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.TranchesLeft))
		i--
		dAtA[i] = 0x38
	}
	if m.Tranches != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Tranches))
		i--
		dAtA[i] = 0x30
	}
	if m.NextClaimFrom != nil {
		n2, err2 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.NextClaimFrom, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.NextClaimFrom):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintTypes(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x2a
	}
	if m.StartTime != nil {
		n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.StartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.StartTime):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintTypes(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x22
	}
	if m.ClaimedAmount != nil {
		{
			size, err := m.ClaimedAmount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.TotalBudget != nil {
		{
			size, err := m.TotalBudget.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.RecipientAddress) > 0 {
		i -= len(m.RecipientAddress)
		copy(dAtA[i:], m.RecipientAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.RecipientAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContinuousFund) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContinuousFund) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContinuousFund) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Expiry != nil {
		n6, err6 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(*m.Expiry, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiry):])
		if err6 != nil {
			return 0, err6
		}
		i -= n6
		i = encodeVarintTypes(dAtA, i, uint64(n6))
		i--
		dAtA[i] = 0x22
	}
	if m.Cap != nil {
		{
			size, err := m.Cap.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	{
		size := m.Percentage.Size()
		i -= size
		if _, err := m.Percentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Recipient) > 0 {
		i -= len(m.Recipient)
		copy(dAtA[i:], m.Recipient)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Recipient)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Budget) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.RecipientAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.TotalBudget != nil {
		l = m.TotalBudget.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.ClaimedAmount != nil {
		l = m.ClaimedAmount.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.StartTime != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.StartTime)
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.NextClaimFrom != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.NextClaimFrom)
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Tranches != 0 {
		n += 1 + sovTypes(uint64(m.Tranches))
	}
	if m.TranchesLeft != 0 {
		n += 1 + sovTypes(uint64(m.TranchesLeft))
	}
	if m.Period != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(*m.Period)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ContinuousFund) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Recipient)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Percentage.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.Cap != nil {
		l = m.Cap.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Expiry != nil {
		l = github_com_cosmos_gogoproto_types.SizeOfStdTime(*m.Expiry)
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Budget) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Budget: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Budget: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecipientAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecipientAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalBudget", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TotalBudget == nil {
				m.TotalBudget = &types.Coin{}
			}
			if err := m.TotalBudget.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClaimedAmount == nil {
				m.ClaimedAmount = &types.Coin{}
			}
			if err := m.ClaimedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartTime == nil {
				m.StartTime = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextClaimFrom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NextClaimFrom == nil {
				m.NextClaimFrom = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.NextClaimFrom, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tranches", wireType)
			}
			m.Tranches = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Tranches |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TranchesLeft", wireType)
			}
			m.TranchesLeft = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TranchesLeft |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Period == nil {
				m.Period = new(time.Duration)
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(m.Period, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ContinuousFund) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ContinuousFund: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContinuousFund: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recipient", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Recipient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Percentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Percentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Cap == nil {
				m.Cap = &types.Coin{}
			}
			if err := m.Cap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Expiry == nil {
				m.Expiry = new(time.Time)
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(m.Expiry, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
