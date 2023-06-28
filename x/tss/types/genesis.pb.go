// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tss/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

// GenesisState defines the tss module's genesis state.
type GenesisState struct {
	// params defines all the paramiters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// group_count defines the number of groups.
	GroupCount uint64 `protobuf:"varint,2,opt,name=group_count,json=groupCount,proto3" json:"group_count,omitempty"`
	// signing_count defines the number of signers.
	SigningCount uint64 `protobuf:"varint,3,opt,name=signing_count,json=signingCount,proto3" json:"signing_count,omitempty"`
	// groups is an array containing information about each group.
	Groups []Group `protobuf:"bytes,4,rep,name=groups,proto3" json:"groups"`
	// de_queues_genesis is an array containing the de queues of all the address.
	DEQueuesGenesis []DEQueueGenesis `protobuf:"bytes,5,rep,name=de_queues_genesis,json=deQueuesGenesis,proto3" json:"de_queues_genesis"`
	// des_genesis is an array containing the des of all the address.
	DEsGenesis []DEGenesis `protobuf:"bytes,6,rep,name=des_genesis,json=desGenesis,proto3" json:"des_genesis"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb5f1c2be1950e47, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetGroupCount() uint64 {
	if m != nil {
		return m.GroupCount
	}
	return 0
}

func (m *GenesisState) GetSigningCount() uint64 {
	if m != nil {
		return m.SigningCount
	}
	return 0
}

func (m *GenesisState) GetGroups() []Group {
	if m != nil {
		return m.Groups
	}
	return nil
}

func (m *GenesisState) GetDEQueuesGenesis() []DEQueueGenesis {
	if m != nil {
		return m.DEQueuesGenesis
	}
	return nil
}

func (m *GenesisState) GetDEsGenesis() []DEGenesis {
	if m != nil {
		return m.DEsGenesis
	}
	return nil
}

// Params defines the set of module parameters.
type Params struct {
	// max_group_size is the maximum of the member capacity of the group.
	MaxGroupSize uint64 `protobuf:"varint,1,opt,name=max_group_size,json=maxGroupSize,proto3" json:"max_group_size,omitempty"`
	// max_d_e_size is the maximum of the de capacity of the member.
	MaxDESize uint64 `protobuf:"varint,2,opt,name=max_d_e_size,json=maxDESize,proto3" json:"max_d_e_size,omitempty"`
	// round_period is the duration period where a member can submit message for the round.
	RoundPeriod time.Duration `protobuf:"bytes,3,opt,name=round_period,json=roundPeriod,proto3,stdduration" json:"round_period"`
	// signing_period is the duration period where a member can submit signing message.
	SigningPeriod time.Duration `protobuf:"bytes,4,opt,name=signing_period,json=signingPeriod,proto3,stdduration" json:"signing_period"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb5f1c2be1950e47, []int{1}
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

func (m *Params) GetMaxGroupSize() uint64 {
	if m != nil {
		return m.MaxGroupSize
	}
	return 0
}

func (m *Params) GetMaxDESize() uint64 {
	if m != nil {
		return m.MaxDESize
	}
	return 0
}

func (m *Params) GetRoundPeriod() time.Duration {
	if m != nil {
		return m.RoundPeriod
	}
	return 0
}

func (m *Params) GetSigningPeriod() time.Duration {
	if m != nil {
		return m.SigningPeriod
	}
	return 0
}

// DEQueueGenesis defines an account address and de queue used in the tss module's
// genesis state.
type DEQueueGenesis struct {
	// address is the address of the de holder.
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// de_queue defines the different de queue this balance holds.
	DEQueue *DEQueue `protobuf:"bytes,2,opt,name=de_queue,json=deQueue,proto3" json:"de_queue,omitempty"`
}

func (m *DEQueueGenesis) Reset()         { *m = DEQueueGenesis{} }
func (m *DEQueueGenesis) String() string { return proto.CompactTextString(m) }
func (*DEQueueGenesis) ProtoMessage()    {}
func (*DEQueueGenesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb5f1c2be1950e47, []int{2}
}
func (m *DEQueueGenesis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DEQueueGenesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DEQueueGenesis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DEQueueGenesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DEQueueGenesis.Merge(m, src)
}
func (m *DEQueueGenesis) XXX_Size() int {
	return m.Size()
}
func (m *DEQueueGenesis) XXX_DiscardUnknown() {
	xxx_messageInfo_DEQueueGenesis.DiscardUnknown(m)
}

var xxx_messageInfo_DEQueueGenesis proto.InternalMessageInfo

func (m *DEQueueGenesis) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *DEQueueGenesis) GetDEQueue() *DEQueue {
	if m != nil {
		return m.DEQueue
	}
	return nil
}

// DEGenesis defines an account address and de pair used in the tss module's
// genesis state.
type DEGenesis struct {
	// address is the address of the de holder.
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// index is the index for store de of the address
	Index uint64 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	// de defines the different de this balance holds.
	DE *DE `protobuf:"bytes,3,opt,name=de,proto3" json:"de,omitempty"`
}

func (m *DEGenesis) Reset()         { *m = DEGenesis{} }
func (m *DEGenesis) String() string { return proto.CompactTextString(m) }
func (*DEGenesis) ProtoMessage()    {}
func (*DEGenesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb5f1c2be1950e47, []int{3}
}
func (m *DEGenesis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DEGenesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DEGenesis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DEGenesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DEGenesis.Merge(m, src)
}
func (m *DEGenesis) XXX_Size() int {
	return m.Size()
}
func (m *DEGenesis) XXX_DiscardUnknown() {
	xxx_messageInfo_DEGenesis.DiscardUnknown(m)
}

var xxx_messageInfo_DEGenesis proto.InternalMessageInfo

func (m *DEGenesis) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *DEGenesis) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *DEGenesis) GetDE() *DE {
	if m != nil {
		return m.DE
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tss.v1beta1.GenesisState")
	proto.RegisterType((*Params)(nil), "tss.v1beta1.Params")
	proto.RegisterType((*DEQueueGenesis)(nil), "tss.v1beta1.DEQueueGenesis")
	proto.RegisterType((*DEGenesis)(nil), "tss.v1beta1.DEGenesis")
}

func init() { proto.RegisterFile("tss/v1beta1/genesis.proto", fileDescriptor_eb5f1c2be1950e47) }

var fileDescriptor_eb5f1c2be1950e47 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xc7, 0xe3, 0x34, 0x75, 0xda, 0x71, 0xbe, 0x44, 0xdf, 0x12, 0x20, 0x2d, 0x92, 0x1d, 0x05,
	0x24, 0x72, 0x40, 0x36, 0x09, 0x57, 0x4e, 0x26, 0x69, 0x25, 0xb8, 0x14, 0xf7, 0xc6, 0xc5, 0x6c,
	0xb2, 0x8b, 0x63, 0xa9, 0xf1, 0x06, 0xef, 0xba, 0x32, 0x7d, 0x0a, 0x8e, 0x3c, 0x52, 0x8f, 0x3d,
	0x72, 0x0a, 0xc8, 0xb9, 0xf0, 0x18, 0xc8, 0xbb, 0xeb, 0x28, 0xa9, 0x10, 0xe2, 0x96, 0x99, 0xf9,
	0xcf, 0x6f, 0xb2, 0xff, 0x19, 0xc3, 0x89, 0xe0, 0xdc, 0xbb, 0x1e, 0xcd, 0xa8, 0xc0, 0x23, 0x2f,
	0xa2, 0x09, 0xe5, 0x31, 0x77, 0x57, 0x29, 0x13, 0x0c, 0x59, 0x82, 0x73, 0x57, 0x97, 0x4e, 0xbb,
	0x11, 0x8b, 0x98, 0xcc, 0x7b, 0xe5, 0x2f, 0x25, 0x39, 0xb5, 0x23, 0xc6, 0xa2, 0x2b, 0xea, 0xc9,
	0x68, 0x96, 0x7d, 0xf2, 0x48, 0x96, 0x62, 0x11, 0xb3, 0x44, 0xd7, 0x1f, 0xee, 0xd2, 0x4b, 0x9c,
	0x4c, 0x0f, 0x7e, 0xd5, 0xa1, 0x75, 0xae, 0x66, 0x5d, 0x0a, 0x2c, 0x28, 0x1a, 0x81, 0xb9, 0xc2,
	0x29, 0x5e, 0xf2, 0x9e, 0xd1, 0x37, 0x86, 0xd6, 0xf8, 0x81, 0xbb, 0x33, 0xdb, 0xbd, 0x90, 0x25,
	0xbf, 0x71, 0xbb, 0x76, 0x6a, 0x81, 0x16, 0x22, 0x07, 0xac, 0x28, 0x65, 0xd9, 0x2a, 0x9c, 0xb3,
	0x2c, 0x11, 0xbd, 0x7a, 0xdf, 0x18, 0x36, 0x02, 0x90, 0xa9, 0x37, 0x65, 0x06, 0x3d, 0x85, 0xff,
	0x78, 0x1c, 0x25, 0x71, 0x12, 0x69, 0xc9, 0x81, 0x94, 0xb4, 0x74, 0x52, 0x89, 0x5e, 0x82, 0x29,
	0x5b, 0x78, 0xaf, 0xd1, 0x3f, 0x18, 0x5a, 0x63, 0xb4, 0x37, 0xf8, 0xbc, 0x2c, 0x55, 0x73, 0x95,
	0x0e, 0x7d, 0x84, 0xff, 0x09, 0x0d, 0x3f, 0x67, 0x34, 0xa3, 0x3c, 0xd4, 0x86, 0xf5, 0x0e, 0x65,
	0xf3, 0x93, 0xbd, 0xe6, 0xc9, 0xf4, 0x7d, 0x29, 0xd2, 0xef, 0xf4, 0x1f, 0x97, 0x94, 0x62, 0xed,
	0x74, 0x74, 0x9e, 0xeb, 0x42, 0xd0, 0x21, 0x74, 0x2f, 0x81, 0xde, 0x81, 0x45, 0x76, 0xd8, 0xa6,
	0x64, 0x3f, 0xba, 0xc7, 0xae, 0xb0, 0x48, 0x63, 0x61, 0x32, 0xdd, 0x12, 0x81, 0x6c, 0x61, 0x83,
	0x8d, 0x01, 0xa6, 0xf2, 0x0f, 0x3d, 0x83, 0xf6, 0x12, 0xe7, 0xa1, 0x72, 0x8d, 0xc7, 0x37, 0x54,
	0x9a, 0xdd, 0x08, 0x5a, 0x4b, 0x9c, 0xcb, 0x97, 0x5e, 0xc6, 0x37, 0x14, 0x39, 0x50, 0xc6, 0x21,
	0x09, 0xa9, 0xd2, 0x28, 0x63, 0x8f, 0x97, 0x38, 0x9f, 0x4c, 0xa5, 0xe0, 0x0c, 0x5a, 0x29, 0xcb,
	0x12, 0x12, 0xae, 0x68, 0x1a, 0x33, 0x22, 0x6d, 0xb5, 0xc6, 0x27, 0xae, 0x3a, 0x05, 0xb7, 0x3a,
	0x05, 0x77, 0xa2, 0x4f, 0xc1, 0x3f, 0x2a, 0xff, 0xe2, 0xb7, 0x1f, 0x8e, 0x11, 0x58, 0xb2, 0xf1,
	0x42, 0xf6, 0xa1, 0xb7, 0xd0, 0xae, 0xf6, 0xa3, 0x49, 0x8d, 0x7f, 0x27, 0x55, 0xab, 0x55, 0xac,
	0xc1, 0x02, 0xda, 0xfb, 0x76, 0xa3, 0x1e, 0x34, 0x31, 0x21, 0x29, 0xe5, 0xea, 0xa4, 0x5a, 0x41,
	0x15, 0xa2, 0xd7, 0x70, 0x54, 0x2d, 0x50, 0x3e, 0xce, 0x1a, 0x77, 0xff, 0xb4, 0x37, 0xdf, 0x2a,
	0xd6, 0x4e, 0x53, 0x07, 0x41, 0x53, 0x2f, 0x69, 0x40, 0xe0, 0x78, 0x6b, 0xfe, 0x5f, 0x86, 0x74,
	0xe1, 0x30, 0x4e, 0x08, 0xcd, 0xb5, 0x7d, 0x2a, 0x40, 0xcf, 0xa1, 0x4e, 0xa8, 0x36, 0xac, 0x73,
	0x6f, 0xa8, 0x6f, 0x16, 0x6b, 0xa7, 0x3e, 0x99, 0x06, 0x75, 0x42, 0xfd, 0xb3, 0xdb, 0xc2, 0x36,
	0xee, 0x0a, 0xdb, 0xf8, 0x59, 0xd8, 0xc6, 0xd7, 0x8d, 0x5d, 0xbb, 0xdb, 0xd8, 0xb5, 0xef, 0x1b,
	0xbb, 0xf6, 0xe1, 0x45, 0x14, 0x8b, 0x45, 0x36, 0x73, 0xe7, 0x6c, 0xe9, 0xcd, 0x70, 0x42, 0xa4,
	0x49, 0x73, 0x76, 0xe5, 0xcd, 0x17, 0x38, 0x4e, 0xbc, 0xeb, 0xb1, 0x97, 0x97, 0x1f, 0x9a, 0x27,
	0xbe, 0xac, 0x28, 0x9f, 0x99, 0xb2, 0xfc, 0xea, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x6f,
	0x00, 0xf3, 0xe6, 0x03, 0x00, 0x00,
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
	if len(m.DEsGenesis) > 0 {
		for iNdEx := len(m.DEsGenesis) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DEsGenesis[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DEQueuesGenesis) > 0 {
		for iNdEx := len(m.DEQueuesGenesis) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DEQueuesGenesis[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Groups) > 0 {
		for iNdEx := len(m.Groups) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Groups[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.SigningCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.SigningCount))
		i--
		dAtA[i] = 0x18
	}
	if m.GroupCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.GroupCount))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
	n2, err2 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.SigningPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.SigningPeriod):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	n3, err3 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.RoundPeriod, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.RoundPeriod):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintGenesis(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x1a
	if m.MaxDESize != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxDESize))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxGroupSize != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.MaxGroupSize))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DEQueueGenesis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DEQueueGenesis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DEQueueGenesis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DEQueue != nil {
		{
			size, err := m.DEQueue.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DEGenesis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DEGenesis) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DEGenesis) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DE != nil {
		{
			size, err := m.DE.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Index != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if m.GroupCount != 0 {
		n += 1 + sovGenesis(uint64(m.GroupCount))
	}
	if m.SigningCount != 0 {
		n += 1 + sovGenesis(uint64(m.SigningCount))
	}
	if len(m.Groups) > 0 {
		for _, e := range m.Groups {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DEQueuesGenesis) > 0 {
		for _, e := range m.DEQueuesGenesis {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DEsGenesis) > 0 {
		for _, e := range m.DEsGenesis {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxGroupSize != 0 {
		n += 1 + sovGenesis(uint64(m.MaxGroupSize))
	}
	if m.MaxDESize != 0 {
		n += 1 + sovGenesis(uint64(m.MaxDESize))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.RoundPeriod)
	n += 1 + l + sovGenesis(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.SigningPeriod)
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *DEQueueGenesis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.DEQueue != nil {
		l = m.DEQueue.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *DEGenesis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Index != 0 {
		n += 1 + sovGenesis(uint64(m.Index))
	}
	if m.DE != nil {
		l = m.DE.Size()
		n += 1 + l + sovGenesis(uint64(l))
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field GroupCount", wireType)
			}
			m.GroupCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.GroupCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningCount", wireType)
			}
			m.SigningCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SigningCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
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
			m.Groups = append(m.Groups, Group{})
			if err := m.Groups[len(m.Groups)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DEQueuesGenesis", wireType)
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
			m.DEQueuesGenesis = append(m.DEQueuesGenesis, DEQueueGenesis{})
			if err := m.DEQueuesGenesis[len(m.DEQueuesGenesis)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DEsGenesis", wireType)
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
			m.DEsGenesis = append(m.DEsGenesis, DEGenesis{})
			if err := m.DEsGenesis[len(m.DEsGenesis)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Params) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxGroupSize", wireType)
			}
			m.MaxGroupSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxGroupSize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxDESize", wireType)
			}
			m.MaxDESize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxDESize |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoundPeriod", wireType)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.RoundPeriod, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SigningPeriod", wireType)
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
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.SigningPeriod, dAtA[iNdEx:postIndex]); err != nil {
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
func (m *DEQueueGenesis) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DEQueueGenesis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DEQueueGenesis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DEQueue", wireType)
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
			if m.DEQueue == nil {
				m.DEQueue = &DEQueue{}
			}
			if err := m.DEQueue.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *DEGenesis) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DEGenesis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DEGenesis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DE", wireType)
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
			if m.DE == nil {
				m.DE = &DE{}
			}
			if err := m.DE.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
