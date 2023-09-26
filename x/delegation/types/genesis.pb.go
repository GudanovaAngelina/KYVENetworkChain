// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kyve/delegation/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the delegation module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// delegator_list ...
	DelegatorList []Delegator `protobuf:"bytes,2,rep,name=delegator_list,json=delegatorList,proto3" json:"delegator_list"`
	// delegation_entry_list ...
	DelegationEntryList []DelegationEntry `protobuf:"bytes,3,rep,name=delegation_entry_list,json=delegationEntryList,proto3" json:"delegation_entry_list"`
	// delegation_data_list ...
	DelegationDataList []DelegationData `protobuf:"bytes,4,rep,name=delegation_data_list,json=delegationDataList,proto3" json:"delegation_data_list"`
	// delegation_slash_list ...
	DelegationSlashList []DelegationSlash `protobuf:"bytes,5,rep,name=delegation_slash_list,json=delegationSlashList,proto3" json:"delegation_slash_list"`
	// undelegation_queue_entry_list ...
	UndelegationQueueEntryList []UndelegationQueueEntry `protobuf:"bytes,6,rep,name=undelegation_queue_entry_list,json=undelegationQueueEntryList,proto3" json:"undelegation_queue_entry_list"`
	// queue_state_undelegation ...
	QueueStateUndelegation QueueState `protobuf:"bytes,7,opt,name=queue_state_undelegation,json=queueStateUndelegation,proto3" json:"queue_state_undelegation"`
	// redelegation_cooldown_list ...
	RedelegationCooldownList []RedelegationCooldown `protobuf:"bytes,8,rep,name=redelegation_cooldown_list,json=redelegationCooldownList,proto3" json:"redelegation_cooldown_list"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bd28fed64b7905b, []int{0}
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

func (m *GenesisState) GetDelegatorList() []Delegator {
	if m != nil {
		return m.DelegatorList
	}
	return nil
}

func (m *GenesisState) GetDelegationEntryList() []DelegationEntry {
	if m != nil {
		return m.DelegationEntryList
	}
	return nil
}

func (m *GenesisState) GetDelegationDataList() []DelegationData {
	if m != nil {
		return m.DelegationDataList
	}
	return nil
}

func (m *GenesisState) GetDelegationSlashList() []DelegationSlash {
	if m != nil {
		return m.DelegationSlashList
	}
	return nil
}

func (m *GenesisState) GetUndelegationQueueEntryList() []UndelegationQueueEntry {
	if m != nil {
		return m.UndelegationQueueEntryList
	}
	return nil
}

func (m *GenesisState) GetQueueStateUndelegation() QueueState {
	if m != nil {
		return m.QueueStateUndelegation
	}
	return QueueState{}
}

func (m *GenesisState) GetRedelegationCooldownList() []RedelegationCooldown {
	if m != nil {
		return m.RedelegationCooldownList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "kyve.delegation.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("kyve/delegation/v1beta1/genesis.proto", fileDescriptor_0bd28fed64b7905b)
}

var fileDescriptor_0bd28fed64b7905b = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x4f, 0xef, 0xd2, 0x30,
	0x1c, 0xc6, 0x37, 0x7f, 0x38, 0x4d, 0x51, 0x0f, 0x13, 0x75, 0x59, 0xe2, 0x20, 0xa8, 0x71, 0x17,
	0xd7, 0x80, 0x67, 0x2f, 0x08, 0x31, 0x46, 0xe3, 0x1f, 0x88, 0x26, 0x7a, 0x59, 0xba, 0xad, 0x19,
	0x0b, 0x63, 0x85, 0xb5, 0xe3, 0xcf, 0xbb, 0xf0, 0xea, 0x3b, 0xe2, 0xc8, 0xd1, 0x93, 0x31, 0xf0,
	0x46, 0xcc, 0xda, 0x0a, 0x25, 0xb2, 0xfc, 0xb8, 0x2d, 0xdf, 0x3d, 0xcf, 0xf3, 0xe9, 0xd3, 0x7c,
	0x0b, 0x9e, 0x4d, 0xd6, 0x0b, 0x0c, 0x23, 0x9c, 0xe2, 0x18, 0xb1, 0x84, 0x64, 0x70, 0xd1, 0x09,
	0x30, 0x43, 0x1d, 0x18, 0xe3, 0x0c, 0xd3, 0x84, 0x7a, 0xb3, 0x9c, 0x30, 0x62, 0x3e, 0x2a, 0x65,
	0xde, 0x51, 0xe6, 0x49, 0x99, 0xdd, 0x88, 0x49, 0x4c, 0xb8, 0x06, 0x96, 0x5f, 0x42, 0x6e, 0xbb,
	0x55, 0xa9, 0x4a, 0x82, 0x50, 0x3e, 0xad, 0x52, 0xce, 0x50, 0x8e, 0xa6, 0x12, 0xdf, 0xfe, 0x69,
	0x80, 0x3b, 0x6f, 0xc4, 0x81, 0x46, 0x0c, 0x31, 0x6c, 0xbe, 0x02, 0x86, 0x10, 0x58, 0x7a, 0x4b,
	0x77, 0xeb, 0xdd, 0xa6, 0x57, 0x71, 0x40, 0xef, 0x13, 0x97, 0xf5, 0x6a, 0x9b, 0xdf, 0x4d, 0x6d,
	0x28, 0x4d, 0xe6, 0x47, 0x70, 0x4f, 0x4a, 0x49, 0xee, 0xa7, 0x09, 0x65, 0xd6, 0x8d, 0xd6, 0x95,
	0x5b, 0xef, 0xb6, 0x2b, 0x63, 0xfa, 0xff, 0xe4, 0x32, 0xe9, 0xee, 0xc1, 0xff, 0x3e, 0xa1, 0xcc,
	0x0c, 0xc0, 0x83, 0xa3, 0xc9, 0xc7, 0x19, 0xcb, 0xd7, 0x22, 0xf7, 0x8a, 0xe7, 0xba, 0xd7, 0xe5,
	0x26, 0x24, 0x1b, 0x94, 0x26, 0x99, 0x7e, 0x3f, 0x3a, 0x1d, 0x73, 0x86, 0x0f, 0x1a, 0x0a, 0x23,
	0x42, 0x0c, 0x09, 0x44, 0x8d, 0x23, 0x9e, 0x5f, 0x80, 0xe8, 0x23, 0x86, 0x24, 0xc1, 0x8c, 0x4e,
	0xa6, 0x67, 0x4a, 0xd0, 0x14, 0xd1, 0xb1, 0x20, 0xdc, 0xbc, 0xb8, 0xc4, 0xa8, 0x34, 0xfd, 0x5f,
	0x82, 0x8f, 0x39, 0x63, 0x05, 0x1e, 0x17, 0x99, 0x42, 0x99, 0x17, 0xb8, 0xc0, 0xea, 0x85, 0x19,
	0x9c, 0x05, 0x2b, 0x59, 0x5f, 0x14, 0xf7, 0xe7, 0xd2, 0xac, 0xde, 0x9b, 0x5d, 0x9c, 0xfd, 0xcb,
	0xc9, 0x21, 0xb0, 0x04, 0x8c, 0x96, 0x1b, 0xe4, 0xab, 0x4a, 0xeb, 0x16, 0x5f, 0xa2, 0x27, 0x95,
	0x50, 0x1e, 0xc5, 0x37, 0x4f, 0x82, 0x1e, 0xce, 0x0f, 0x13, 0xf5, 0x40, 0xe6, 0x1c, 0xd8, 0x39,
	0x56, 0xea, 0x85, 0x84, 0xa4, 0x11, 0x59, 0x66, 0xa2, 0xdb, 0x6d, 0xde, 0xed, 0x45, 0x25, 0x66,
	0xa8, 0x58, 0x5f, 0x4b, 0xa7, 0x04, 0x5a, 0xf9, 0x99, 0x7f, 0x65, 0xaf, 0xde, 0xdb, 0xcd, 0xce,
	0xd1, 0xb7, 0x3b, 0x47, 0xff, 0xb3, 0x73, 0xf4, 0x1f, 0x7b, 0x47, 0xdb, 0xee, 0x1d, 0xed, 0xd7,
	0xde, 0xd1, 0xbe, 0xc3, 0x38, 0x61, 0xe3, 0x22, 0xf0, 0x42, 0x32, 0x85, 0xef, 0xbe, 0x7d, 0x1d,
	0x7c, 0xc0, 0x6c, 0x49, 0xf2, 0x09, 0x0c, 0xc7, 0x28, 0xc9, 0xe0, 0x4a, 0x7d, 0x75, 0x6c, 0x3d,
	0xc3, 0x34, 0x30, 0xf8, 0x6b, 0x7b, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa8, 0xa4, 0xfe, 0x14,
	0x15, 0x04, 0x00, 0x00,
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
	if len(m.RedelegationCooldownList) > 0 {
		for iNdEx := len(m.RedelegationCooldownList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RedelegationCooldownList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	{
		size, err := m.QueueStateUndelegation.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	if len(m.UndelegationQueueEntryList) > 0 {
		for iNdEx := len(m.UndelegationQueueEntryList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UndelegationQueueEntryList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DelegationSlashList) > 0 {
		for iNdEx := len(m.DelegationSlashList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationSlashList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DelegationDataList) > 0 {
		for iNdEx := len(m.DelegationDataList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationDataList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DelegationEntryList) > 0 {
		for iNdEx := len(m.DelegationEntryList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegationEntryList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DelegatorList) > 0 {
		for iNdEx := len(m.DelegatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DelegatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DelegatorList) > 0 {
		for _, e := range m.DelegatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DelegationEntryList) > 0 {
		for _, e := range m.DelegationEntryList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DelegationDataList) > 0 {
		for _, e := range m.DelegationDataList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DelegationSlashList) > 0 {
		for _, e := range m.DelegationSlashList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UndelegationQueueEntryList) > 0 {
		for _, e := range m.UndelegationQueueEntryList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.QueueStateUndelegation.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.RedelegationCooldownList) > 0 {
		for _, e := range m.RedelegationCooldownList {
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegatorList", wireType)
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
			m.DelegatorList = append(m.DelegatorList, Delegator{})
			if err := m.DelegatorList[len(m.DelegatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationEntryList", wireType)
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
			m.DelegationEntryList = append(m.DelegationEntryList, DelegationEntry{})
			if err := m.DelegationEntryList[len(m.DelegationEntryList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationDataList", wireType)
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
			m.DelegationDataList = append(m.DelegationDataList, DelegationData{})
			if err := m.DelegationDataList[len(m.DelegationDataList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DelegationSlashList", wireType)
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
			m.DelegationSlashList = append(m.DelegationSlashList, DelegationSlash{})
			if err := m.DelegationSlashList[len(m.DelegationSlashList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UndelegationQueueEntryList", wireType)
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
			m.UndelegationQueueEntryList = append(m.UndelegationQueueEntryList, UndelegationQueueEntry{})
			if err := m.UndelegationQueueEntryList[len(m.UndelegationQueueEntryList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueueStateUndelegation", wireType)
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
			if err := m.QueueStateUndelegation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RedelegationCooldownList", wireType)
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
			m.RedelegationCooldownList = append(m.RedelegationCooldownList, RedelegationCooldown{})
			if err := m.RedelegationCooldownList[len(m.RedelegationCooldownList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
