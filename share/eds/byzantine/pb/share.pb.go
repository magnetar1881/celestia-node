// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: share/eds/byzantine/pb/share.proto

package pb

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
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

type Axis int32

const (
	Axis_ROW Axis = 0
	Axis_COL Axis = 1
)

var Axis_name = map[int32]string{
	0: "ROW",
	1: "COL",
}

var Axis_value = map[string]int32{
	"ROW": 0,
	"COL": 1,
}

func (x Axis) String() string {
	return proto.EnumName(Axis_name, int32(x))
}

func (Axis) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d28ce8f160a920d1, []int{0}
}

type MerkleProof struct {
	Start    int64    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End      int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	Nodes    [][]byte `protobuf:"bytes,3,rep,name=nodes,proto3" json:"nodes,omitempty"`
	LeafHash []byte   `protobuf:"bytes,4,opt,name=leaf_hash,json=leafHash,proto3" json:"leaf_hash,omitempty"`
}

func (m *MerkleProof) Reset()         { *m = MerkleProof{} }
func (m *MerkleProof) String() string { return proto.CompactTextString(m) }
func (*MerkleProof) ProtoMessage()    {}
func (*MerkleProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28ce8f160a920d1, []int{0}
}
func (m *MerkleProof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MerkleProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MerkleProof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MerkleProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerkleProof.Merge(m, src)
}
func (m *MerkleProof) XXX_Size() int {
	return m.Size()
}
func (m *MerkleProof) XXX_DiscardUnknown() {
	xxx_messageInfo_MerkleProof.DiscardUnknown(m)
}

var xxx_messageInfo_MerkleProof proto.InternalMessageInfo

func (m *MerkleProof) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *MerkleProof) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

func (m *MerkleProof) GetNodes() [][]byte {
	if m != nil {
		return m.Nodes
	}
	return nil
}

func (m *MerkleProof) GetLeafHash() []byte {
	if m != nil {
		return m.LeafHash
	}
	return nil
}

type Share struct {
	Data  []byte       `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
	Proof *MerkleProof `protobuf:"bytes,2,opt,name=Proof,proto3" json:"Proof,omitempty"`
}

func (m *Share) Reset()         { *m = Share{} }
func (m *Share) String() string { return proto.CompactTextString(m) }
func (*Share) ProtoMessage()    {}
func (*Share) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28ce8f160a920d1, []int{1}
}
func (m *Share) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Share) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Share.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Share) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Share.Merge(m, src)
}
func (m *Share) XXX_Size() int {
	return m.Size()
}
func (m *Share) XXX_DiscardUnknown() {
	xxx_messageInfo_Share.DiscardUnknown(m)
}

var xxx_messageInfo_Share proto.InternalMessageInfo

func (m *Share) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Share) GetProof() *MerkleProof {
	if m != nil {
		return m.Proof
	}
	return nil
}

type BadEncoding struct {
	HeaderHash []byte   `protobuf:"bytes,1,opt,name=HeaderHash,proto3" json:"HeaderHash,omitempty"`
	Height     uint64   `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	Shares     []*Share `protobuf:"bytes,3,rep,name=Shares,proto3" json:"Shares,omitempty"`
	Index      uint32   `protobuf:"varint,4,opt,name=Index,proto3" json:"Index,omitempty"`
	Axis       Axis     `protobuf:"varint,5,opt,name=Axis,proto3,enum=share.eds.byzantine.pb.Axis" json:"Axis,omitempty"`
}

func (m *BadEncoding) Reset()         { *m = BadEncoding{} }
func (m *BadEncoding) String() string { return proto.CompactTextString(m) }
func (*BadEncoding) ProtoMessage()    {}
func (*BadEncoding) Descriptor() ([]byte, []int) {
	return fileDescriptor_d28ce8f160a920d1, []int{2}
}
func (m *BadEncoding) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BadEncoding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BadEncoding.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BadEncoding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadEncoding.Merge(m, src)
}
func (m *BadEncoding) XXX_Size() int {
	return m.Size()
}
func (m *BadEncoding) XXX_DiscardUnknown() {
	xxx_messageInfo_BadEncoding.DiscardUnknown(m)
}

var xxx_messageInfo_BadEncoding proto.InternalMessageInfo

func (m *BadEncoding) GetHeaderHash() []byte {
	if m != nil {
		return m.HeaderHash
	}
	return nil
}

func (m *BadEncoding) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BadEncoding) GetShares() []*Share {
	if m != nil {
		return m.Shares
	}
	return nil
}

func (m *BadEncoding) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BadEncoding) GetAxis() Axis {
	if m != nil {
		return m.Axis
	}
	return Axis_ROW
}

func init() {
	proto.RegisterEnum("share.eds.byzantine.pb.Axis", Axis_name, Axis_value)
	proto.RegisterType((*MerkleProof)(nil), "share.eds.byzantine.pb.MerkleProof")
	proto.RegisterType((*Share)(nil), "share.eds.byzantine.pb.Share")
	proto.RegisterType((*BadEncoding)(nil), "share.eds.byzantine.pb.BadEncoding")
}

func init() {
	proto.RegisterFile("share/eds/byzantine/pb/share.proto", fileDescriptor_d28ce8f160a920d1)
}

var fileDescriptor_d28ce8f160a920d1 = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x41, 0x6b, 0xdb, 0x30,
	0x18, 0xb5, 0x66, 0x3b, 0xdb, 0x3e, 0x67, 0x23, 0x88, 0x11, 0x0c, 0xdb, 0x8c, 0xf1, 0x2e, 0x66,
	0x30, 0x7b, 0x64, 0xec, 0x30, 0x76, 0x6a, 0xda, 0x42, 0x0a, 0x2d, 0x29, 0x2a, 0x6d, 0xa1, 0x97,
	0x22, 0x47, 0x8a, 0x6d, 0x9a, 0x5a, 0xc1, 0x52, 0x21, 0xed, 0xaf, 0xe8, 0x8f, 0xea, 0xa1, 0xc7,
	0x1c, 0x7b, 0x2c, 0xc9, 0x1f, 0x29, 0x92, 0x43, 0xc8, 0x21, 0xb9, 0xbd, 0xf7, 0x78, 0xd2, 0x7b,
	0x9f, 0xf4, 0x41, 0x24, 0x0b, 0x5a, 0xf3, 0x94, 0x33, 0x99, 0x66, 0xf7, 0x0f, 0xb4, 0x52, 0x65,
	0xc5, 0xd3, 0x69, 0x96, 0x1a, 0x39, 0x99, 0xd6, 0x42, 0x09, 0xdc, 0x6d, 0x08, 0x67, 0x32, 0x59,
	0x7b, 0x92, 0x69, 0x16, 0x15, 0xe0, 0x9d, 0xf0, 0xfa, 0x66, 0xc2, 0x4f, 0x6b, 0x21, 0xc6, 0xf8,
	0x0b, 0xb8, 0x52, 0xd1, 0x5a, 0xf9, 0x28, 0x44, 0xb1, 0x4d, 0x1a, 0x82, 0x3b, 0x60, 0xf3, 0x8a,
	0xf9, 0xef, 0x8c, 0xa6, 0xa1, 0xf6, 0x55, 0x82, 0x71, 0xe9, 0xdb, 0xa1, 0x1d, 0xb7, 0x49, 0x43,
	0xf0, 0x57, 0xf8, 0x38, 0xe1, 0x74, 0x7c, 0x5d, 0x50, 0x59, 0xf8, 0x4e, 0x88, 0xe2, 0x36, 0xf9,
	0xa0, 0x85, 0x01, 0x95, 0x45, 0x74, 0x01, 0xee, 0x99, 0xee, 0x80, 0x31, 0x38, 0x07, 0x54, 0x51,
	0x13, 0xd1, 0x26, 0x06, 0xe3, 0x7f, 0xe0, 0x9a, 0x02, 0x26, 0xc3, 0xeb, 0xfd, 0x48, 0xb6, 0xd7,
	0x4d, 0x36, 0xba, 0x92, 0xe6, 0x44, 0xf4, 0x84, 0xc0, 0xeb, 0x53, 0x76, 0x58, 0x8d, 0x04, 0x2b,
	0xab, 0x1c, 0x07, 0x00, 0x03, 0x4e, 0x19, 0xaf, 0x75, 0xea, 0x2a, 0x64, 0x43, 0xc1, 0x5d, 0x68,
	0x0d, 0x78, 0x99, 0x17, 0xca, 0x64, 0x39, 0x64, 0xc5, 0xf0, 0x5f, 0x68, 0x99, 0x7e, 0xcd, 0x4c,
	0x5e, 0xef, 0xfb, 0xae, 0x0e, 0xc6, 0x45, 0x56, 0x66, 0xfd, 0x12, 0x47, 0x15, 0xe3, 0x33, 0x33,
	0xef, 0x27, 0xd2, 0x10, 0xfc, 0x1b, 0x9c, 0xbd, 0x59, 0x29, 0x7d, 0x37, 0x44, 0xf1, 0xe7, 0xde,
	0xb7, 0x5d, 0x57, 0xd1, 0x59, 0x29, 0x89, 0x71, 0xfe, 0xf4, 0xc1, 0xd1, 0x0c, 0xbf, 0x07, 0x9b,
	0x0c, 0x2f, 0x3b, 0x96, 0x06, 0xfb, 0xc3, 0xe3, 0x0e, 0xea, 0x9f, 0x3f, 0x2f, 0x02, 0x34, 0x5f,
	0x04, 0xe8, 0x75, 0x11, 0xa0, 0xc7, 0x65, 0x60, 0xcd, 0x97, 0x81, 0xf5, 0xb2, 0x0c, 0xac, 0xab,
	0xff, 0x79, 0xa9, 0x8a, 0xbb, 0x2c, 0x19, 0x89, 0xdb, 0x74, 0xc4, 0x27, 0x5c, 0xaa, 0x92, 0x8a,
	0x3a, 0x5f, 0xe3, 0x5f, 0xfa, 0x5b, 0xd2, 0xed, 0xdb, 0x91, 0xb5, 0xcc, 0x62, 0xfc, 0x79, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x2d, 0xf3, 0x01, 0x31, 0x3e, 0x02, 0x00, 0x00,
}

func (m *MerkleProof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MerkleProof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MerkleProof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LeafHash) > 0 {
		i -= len(m.LeafHash)
		copy(dAtA[i:], m.LeafHash)
		i = encodeVarintShare(dAtA, i, uint64(len(m.LeafHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Nodes) > 0 {
		for iNdEx := len(m.Nodes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Nodes[iNdEx])
			copy(dAtA[i:], m.Nodes[iNdEx])
			i = encodeVarintShare(dAtA, i, uint64(len(m.Nodes[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.End != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.End))
		i--
		dAtA[i] = 0x10
	}
	if m.Start != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.Start))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Share) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Share) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Share) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintShare(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintShare(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BadEncoding) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BadEncoding) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BadEncoding) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Axis != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.Axis))
		i--
		dAtA[i] = 0x28
	}
	if m.Index != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Shares) > 0 {
		for iNdEx := len(m.Shares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Shares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintShare(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Height != 0 {
		i = encodeVarintShare(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.HeaderHash) > 0 {
		i -= len(m.HeaderHash)
		copy(dAtA[i:], m.HeaderHash)
		i = encodeVarintShare(dAtA, i, uint64(len(m.HeaderHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintShare(dAtA []byte, offset int, v uint64) int {
	offset -= sovShare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MerkleProof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Start != 0 {
		n += 1 + sovShare(uint64(m.Start))
	}
	if m.End != 0 {
		n += 1 + sovShare(uint64(m.End))
	}
	if len(m.Nodes) > 0 {
		for _, b := range m.Nodes {
			l = len(b)
			n += 1 + l + sovShare(uint64(l))
		}
	}
	l = len(m.LeafHash)
	if l > 0 {
		n += 1 + l + sovShare(uint64(l))
	}
	return n
}

func (m *Share) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovShare(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovShare(uint64(l))
	}
	return n
}

func (m *BadEncoding) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HeaderHash)
	if l > 0 {
		n += 1 + l + sovShare(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovShare(uint64(m.Height))
	}
	if len(m.Shares) > 0 {
		for _, e := range m.Shares {
			l = e.Size()
			n += 1 + l + sovShare(uint64(l))
		}
	}
	if m.Index != 0 {
		n += 1 + sovShare(uint64(m.Index))
	}
	if m.Axis != 0 {
		n += 1 + sovShare(uint64(m.Axis))
	}
	return n
}

func sovShare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozShare(x uint64) (n int) {
	return sovShare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MerkleProof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
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
			return fmt.Errorf("proto: MerkleProof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MerkleProof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			m.Start = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Start |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			m.End = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.End |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, make([]byte, postIndex-iNdEx))
			copy(m.Nodes[len(m.Nodes)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeafHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LeafHash = append(m.LeafHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LeafHash == nil {
				m.LeafHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func (m *Share) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
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
			return fmt.Errorf("proto: Share: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Share: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &MerkleProof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func (m *BadEncoding) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowShare
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
			return fmt.Errorf("proto: BadEncoding: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BadEncoding: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeaderHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeaderHash = append(m.HeaderHash[:0], dAtA[iNdEx:postIndex]...)
			if m.HeaderHash == nil {
				m.HeaderHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
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
				return ErrInvalidLengthShare
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shares = append(m.Shares, &Share{})
			if err := m.Shares[len(m.Shares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Axis", wireType)
			}
			m.Axis = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Axis |= Axis(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthShare
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
func skipShare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowShare
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
					return 0, ErrIntOverflowShare
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
					return 0, ErrIntOverflowShare
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
				return 0, ErrInvalidLengthShare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupShare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthShare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthShare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowShare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupShare = fmt.Errorf("proto: unexpected end of group")
)