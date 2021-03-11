// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vncarbon/vncarbon.proto

package types

import (
	fmt "fmt"
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

type Vncarbon struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Carbon  int32  `protobuf:"varint,3,opt,name=carbon,proto3" json:"carbon,omitempty"`
	Emitter string `protobuf:"bytes,4,opt,name=emitter,proto3" json:"emitter,omitempty"`
	Date    string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (m *Vncarbon) Reset()         { *m = Vncarbon{} }
func (m *Vncarbon) String() string { return proto.CompactTextString(m) }
func (*Vncarbon) ProtoMessage()    {}
func (*Vncarbon) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ee54cee273e8ffd, []int{0}
}
func (m *Vncarbon) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vncarbon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vncarbon.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vncarbon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vncarbon.Merge(m, src)
}
func (m *Vncarbon) XXX_Size() int {
	return m.Size()
}
func (m *Vncarbon) XXX_DiscardUnknown() {
	xxx_messageInfo_Vncarbon.DiscardUnknown(m)
}

var xxx_messageInfo_Vncarbon proto.InternalMessageInfo

func (m *Vncarbon) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Vncarbon) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vncarbon) GetCarbon() int32 {
	if m != nil {
		return m.Carbon
	}
	return 0
}

func (m *Vncarbon) GetEmitter() string {
	if m != nil {
		return m.Emitter
	}
	return ""
}

func (m *Vncarbon) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type MsgCreateVncarbon struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Carbon  int32  `protobuf:"varint,2,opt,name=carbon,proto3" json:"carbon,omitempty"`
	Emitter string `protobuf:"bytes,3,opt,name=emitter,proto3" json:"emitter,omitempty"`
	Date    string `protobuf:"bytes,4,opt,name=date,proto3" json:"date,omitempty"`
}

func (m *MsgCreateVncarbon) Reset()         { *m = MsgCreateVncarbon{} }
func (m *MsgCreateVncarbon) String() string { return proto.CompactTextString(m) }
func (*MsgCreateVncarbon) ProtoMessage()    {}
func (*MsgCreateVncarbon) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ee54cee273e8ffd, []int{1}
}
func (m *MsgCreateVncarbon) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateVncarbon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateVncarbon.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateVncarbon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateVncarbon.Merge(m, src)
}
func (m *MsgCreateVncarbon) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateVncarbon) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateVncarbon.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateVncarbon proto.InternalMessageInfo

func (m *MsgCreateVncarbon) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateVncarbon) GetCarbon() int32 {
	if m != nil {
		return m.Carbon
	}
	return 0
}

func (m *MsgCreateVncarbon) GetEmitter() string {
	if m != nil {
		return m.Emitter
	}
	return ""
}

func (m *MsgCreateVncarbon) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type MsgUpdateVncarbon struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Carbon  int32  `protobuf:"varint,3,opt,name=carbon,proto3" json:"carbon,omitempty"`
	Emitter string `protobuf:"bytes,4,opt,name=emitter,proto3" json:"emitter,omitempty"`
	Date    string `protobuf:"bytes,5,opt,name=date,proto3" json:"date,omitempty"`
}

func (m *MsgUpdateVncarbon) Reset()         { *m = MsgUpdateVncarbon{} }
func (m *MsgUpdateVncarbon) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateVncarbon) ProtoMessage()    {}
func (*MsgUpdateVncarbon) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ee54cee273e8ffd, []int{2}
}
func (m *MsgUpdateVncarbon) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateVncarbon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateVncarbon.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateVncarbon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateVncarbon.Merge(m, src)
}
func (m *MsgUpdateVncarbon) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateVncarbon) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateVncarbon.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateVncarbon proto.InternalMessageInfo

func (m *MsgUpdateVncarbon) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateVncarbon) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MsgUpdateVncarbon) GetCarbon() int32 {
	if m != nil {
		return m.Carbon
	}
	return 0
}

func (m *MsgUpdateVncarbon) GetEmitter() string {
	if m != nil {
		return m.Emitter
	}
	return ""
}

func (m *MsgUpdateVncarbon) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

type MsgDeleteVncarbon struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgDeleteVncarbon) Reset()         { *m = MsgDeleteVncarbon{} }
func (m *MsgDeleteVncarbon) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteVncarbon) ProtoMessage()    {}
func (*MsgDeleteVncarbon) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ee54cee273e8ffd, []int{3}
}
func (m *MsgDeleteVncarbon) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteVncarbon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteVncarbon.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteVncarbon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteVncarbon.Merge(m, src)
}
func (m *MsgDeleteVncarbon) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteVncarbon) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteVncarbon.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteVncarbon proto.InternalMessageInfo

func (m *MsgDeleteVncarbon) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeleteVncarbon) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Vncarbon)(nil), "faddat.vncarbon.vncarbon.Vncarbon")
	proto.RegisterType((*MsgCreateVncarbon)(nil), "faddat.vncarbon.vncarbon.MsgCreateVncarbon")
	proto.RegisterType((*MsgUpdateVncarbon)(nil), "faddat.vncarbon.vncarbon.MsgUpdateVncarbon")
	proto.RegisterType((*MsgDeleteVncarbon)(nil), "faddat.vncarbon.vncarbon.MsgDeleteVncarbon")
}

func init() { proto.RegisterFile("vncarbon/vncarbon.proto", fileDescriptor_7ee54cee273e8ffd) }

var fileDescriptor_7ee54cee273e8ffd = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0xcb, 0x4b, 0x4e,
	0x2c, 0x4a, 0xca, 0xcf, 0xd3, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x24, 0xd2,
	0x12, 0x53, 0x52, 0x12, 0x4b, 0xf4, 0xe0, 0xc2, 0x30, 0x86, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e,
	0x58, 0x91, 0x3e, 0x88, 0x05, 0x51, 0xaf, 0x54, 0xc5, 0xc5, 0x11, 0x06, 0x55, 0x21, 0x24, 0xc1,
	0xc5, 0x9e, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0x5f, 0x24, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04,
	0xe3, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x81, 0x05, 0x99, 0x32, 0x53, 0x84, 0xc4,
	0xb8, 0xd8, 0x20, 0x7a, 0x24, 0x98, 0x15, 0x18, 0x35, 0x58, 0x83, 0xd8, 0x10, 0x26, 0xa4, 0xe6,
	0x66, 0x96, 0x94, 0xa4, 0x16, 0x49, 0xb0, 0x40, 0x4c, 0x80, 0x72, 0x85, 0x84, 0xb8, 0x58, 0x52,
	0x12, 0x4b, 0x52, 0x25, 0x58, 0xc1, 0xc2, 0x60, 0xb6, 0x52, 0x31, 0x97, 0xa0, 0x6f, 0x71, 0xba,
	0x33, 0xc8, 0x8e, 0x54, 0x22, 0x1c, 0x81, 0xb0, 0x94, 0x09, 0x97, 0xa5, 0xcc, 0xd8, 0x2d, 0x65,
	0x41, 0xb2, 0xb4, 0x99, 0x11, 0x6c, 0x6b, 0x68, 0x41, 0x0a, 0x71, 0xb6, 0xd2, 0xc6, 0xeb, 0xb6,
	0x60, 0x47, 0xb8, 0xa4, 0xe6, 0xa4, 0x92, 0xe3, 0x08, 0x27, 0xd7, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18,
	0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf,
	0xd5, 0x87, 0x24, 0x05, 0x78, 0x0a, 0xd1, 0xaf, 0x40, 0x30, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93,
	0xd8, 0xc0, 0x69, 0xc0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xf2, 0xb7, 0xc6, 0x4e, 0x02,
	0x00, 0x00,
}

func (m *Vncarbon) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vncarbon) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vncarbon) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Emitter) > 0 {
		i -= len(m.Emitter)
		copy(dAtA[i:], m.Emitter)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Emitter)))
		i--
		dAtA[i] = 0x22
	}
	if m.Carbon != 0 {
		i = encodeVarintVncarbon(dAtA, i, uint64(m.Carbon))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateVncarbon) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateVncarbon) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateVncarbon) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Emitter) > 0 {
		i -= len(m.Emitter)
		copy(dAtA[i:], m.Emitter)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Emitter)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Carbon != 0 {
		i = encodeVarintVncarbon(dAtA, i, uint64(m.Carbon))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateVncarbon) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateVncarbon) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateVncarbon) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Emitter) > 0 {
		i -= len(m.Emitter)
		copy(dAtA[i:], m.Emitter)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Emitter)))
		i--
		dAtA[i] = 0x22
	}
	if m.Carbon != 0 {
		i = encodeVarintVncarbon(dAtA, i, uint64(m.Carbon))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteVncarbon) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteVncarbon) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteVncarbon) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintVncarbon(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVncarbon(dAtA []byte, offset int, v uint64) int {
	offset -= sovVncarbon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Vncarbon) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	if m.Carbon != 0 {
		n += 1 + sovVncarbon(uint64(m.Carbon))
	}
	l = len(m.Emitter)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	return n
}

func (m *MsgCreateVncarbon) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	if m.Carbon != 0 {
		n += 1 + sovVncarbon(uint64(m.Carbon))
	}
	l = len(m.Emitter)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	return n
}

func (m *MsgUpdateVncarbon) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	if m.Carbon != 0 {
		n += 1 + sovVncarbon(uint64(m.Carbon))
	}
	l = len(m.Emitter)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	return n
}

func (m *MsgDeleteVncarbon) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovVncarbon(uint64(l))
	}
	return n
}

func sovVncarbon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVncarbon(x uint64) (n int) {
	return sovVncarbon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vncarbon) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVncarbon
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
			return fmt.Errorf("proto: Vncarbon: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vncarbon: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Carbon", wireType)
			}
			m.Carbon = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Carbon |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Emitter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Emitter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVncarbon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVncarbon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVncarbon
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
func (m *MsgCreateVncarbon) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVncarbon
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
			return fmt.Errorf("proto: MsgCreateVncarbon: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateVncarbon: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Carbon", wireType)
			}
			m.Carbon = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Carbon |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Emitter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Emitter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVncarbon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVncarbon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVncarbon
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
func (m *MsgUpdateVncarbon) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVncarbon
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
			return fmt.Errorf("proto: MsgUpdateVncarbon: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateVncarbon: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Carbon", wireType)
			}
			m.Carbon = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Carbon |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Emitter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Emitter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVncarbon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVncarbon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVncarbon
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
func (m *MsgDeleteVncarbon) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVncarbon
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
			return fmt.Errorf("proto: MsgDeleteVncarbon: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteVncarbon: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVncarbon
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
				return ErrInvalidLengthVncarbon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVncarbon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVncarbon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthVncarbon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthVncarbon
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
func skipVncarbon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVncarbon
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
					return 0, ErrIntOverflowVncarbon
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
					return 0, ErrIntOverflowVncarbon
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
				return 0, ErrInvalidLengthVncarbon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVncarbon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVncarbon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVncarbon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVncarbon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVncarbon = fmt.Errorf("proto: unexpected end of group")
)