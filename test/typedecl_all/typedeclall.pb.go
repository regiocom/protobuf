// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: typedeclall.proto

package typedeclall

import (
	fmt "fmt"
	_ "github.com/regiocom/protobuf/gogoproto"
	proto "github.com/regiocom/protobuf/proto"
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

func (m *Dropped) Reset()         { *m = Dropped{} }
func (m *Dropped) String() string { return proto.CompactTextString(m) }
func (*Dropped) ProtoMessage()    {}
func (*Dropped) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18a9e6d5cf09c5a, []int{0}
}
func (m *Dropped) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Dropped) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Dropped.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Dropped) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dropped.Merge(m, src)
}
func (m *Dropped) XXX_Size() int {
	return m.Size()
}
func (m *Dropped) XXX_DiscardUnknown() {
	xxx_messageInfo_Dropped.DiscardUnknown(m)
}

var xxx_messageInfo_Dropped proto.InternalMessageInfo

func (m *Dropped) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dropped) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *DroppedWithoutGetters) Reset()         { *m = DroppedWithoutGetters{} }
func (m *DroppedWithoutGetters) String() string { return proto.CompactTextString(m) }
func (*DroppedWithoutGetters) ProtoMessage()    {}
func (*DroppedWithoutGetters) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18a9e6d5cf09c5a, []int{1}
}
func (m *DroppedWithoutGetters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DroppedWithoutGetters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DroppedWithoutGetters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DroppedWithoutGetters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DroppedWithoutGetters.Merge(m, src)
}
func (m *DroppedWithoutGetters) XXX_Size() int {
	return m.Size()
}
func (m *DroppedWithoutGetters) XXX_DiscardUnknown() {
	xxx_messageInfo_DroppedWithoutGetters.DiscardUnknown(m)
}

var xxx_messageInfo_DroppedWithoutGetters proto.InternalMessageInfo

type Kept struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Kept) Reset()         { *m = Kept{} }
func (m *Kept) String() string { return proto.CompactTextString(m) }
func (*Kept) ProtoMessage()    {}
func (*Kept) Descriptor() ([]byte, []int) {
	return fileDescriptor_a18a9e6d5cf09c5a, []int{2}
}
func (m *Kept) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Kept) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Kept.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Kept) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Kept.Merge(m, src)
}
func (m *Kept) XXX_Size() int {
	return m.Size()
}
func (m *Kept) XXX_DiscardUnknown() {
	xxx_messageInfo_Kept.DiscardUnknown(m)
}

var xxx_messageInfo_Kept proto.InternalMessageInfo

func (m *Kept) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Kept) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*Dropped)(nil), "typedeclall.Dropped")
	proto.RegisterType((*DroppedWithoutGetters)(nil), "typedeclall.DroppedWithoutGetters")
	proto.RegisterType((*Kept)(nil), "typedeclall.Kept")
}

func init() { proto.RegisterFile("typedeclall.proto", fileDescriptor_a18a9e6d5cf09c5a) }

var fileDescriptor_a18a9e6d5cf09c5a = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0xa9, 0x2c, 0x48,
	0x4d, 0x49, 0x4d, 0xce, 0x49, 0xcc, 0xc9, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf,
	0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x49, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2, 0x57,
	0x49, 0x9f, 0x8b, 0xdd, 0xa5, 0x28, 0xbf, 0xa0, 0x20, 0x35, 0x45, 0x48, 0x88, 0x8b, 0x25, 0x2f,
	0x31, 0x37, 0x55, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0x12, 0xe0, 0x62, 0x4e,
	0x4c, 0x4f, 0x95, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x0d, 0x02, 0x31, 0x95, 0xbc, 0xb9, 0x44, 0xa1,
	0x1a, 0xc2, 0x33, 0x4b, 0x32, 0xf2, 0x4b, 0x4b, 0xdc, 0x53, 0x4b, 0x4a, 0x52, 0x8b, 0x8a, 0x85,
	0xc4, 0xb8, 0xd8, 0x32, 0x52, 0x33, 0xd3, 0x33, 0x4a, 0xc0, 0x06, 0x30, 0x07, 0x41, 0x79, 0x42,
	0x22, 0x5c, 0xac, 0xe5, 0x99, 0x29, 0x25, 0x19, 0x60, 0x43, 0x98, 0x83, 0x20, 0x1c, 0x2b, 0x96,
	0x8e, 0x05, 0xf2, 0x0c, 0x4a, 0x46, 0x5c, 0x2c, 0xde, 0xa9, 0x05, 0x25, 0xc4, 0x59, 0x6d, 0xc5,
	0xf2, 0x61, 0xa1, 0x3c, 0xa3, 0x93, 0xc1, 0x83, 0x87, 0x72, 0x8c, 0x3f, 0x1e, 0xca, 0x31, 0xae,
	0x78, 0x24, 0xc7, 0xb8, 0xe3, 0x91, 0x1c, 0xe3, 0x81, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8f, 0x47, 0x72, 0x0c, 0x0d, 0x8f, 0xe5,
	0x18, 0x26, 0x3c, 0x96, 0x63, 0xd8, 0xf0, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xec, 0x55, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xed, 0x3d, 0x95, 0x3b, 0x01, 0x00, 0x00,
}

func (this *Dropped) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Dropped)
	if !ok {
		that2, ok := that.(Dropped)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Dropped")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Dropped but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Dropped but is not nil && this == nil")
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.Age != that1.Age {
		return fmt.Errorf("Age this(%v) Not Equal that(%v)", this.Age, that1.Age)
	}
	return nil
}
func (this *Dropped) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Dropped)
	if !ok {
		that2, ok := that.(Dropped)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Age != that1.Age {
		return false
	}
	return true
}
func (this *DroppedWithoutGetters) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*DroppedWithoutGetters)
	if !ok {
		that2, ok := that.(DroppedWithoutGetters)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *DroppedWithoutGetters")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *DroppedWithoutGetters but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *DroppedWithoutGetters but is not nil && this == nil")
	}
	if this.Height != that1.Height {
		return fmt.Errorf("Height this(%v) Not Equal that(%v)", this.Height, that1.Height)
	}
	if this.Width != that1.Width {
		return fmt.Errorf("Width this(%v) Not Equal that(%v)", this.Width, that1.Width)
	}
	return nil
}
func (this *DroppedWithoutGetters) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DroppedWithoutGetters)
	if !ok {
		that2, ok := that.(DroppedWithoutGetters)
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
	if this.Height != that1.Height {
		return false
	}
	if this.Width != that1.Width {
		return false
	}
	return true
}
func (this *Kept) VerboseEqual(that interface{}) error {
	if that == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that == nil && this != nil")
	}

	that1, ok := that.(*Kept)
	if !ok {
		that2, ok := that.(Kept)
		if ok {
			that1 = &that2
		} else {
			return fmt.Errorf("that is not of type *Kept")
		}
	}
	if that1 == nil {
		if this == nil {
			return nil
		}
		return fmt.Errorf("that is type *Kept but is nil && this != nil")
	} else if this == nil {
		return fmt.Errorf("that is type *Kept but is not nil && this == nil")
	}
	if this.Name != that1.Name {
		return fmt.Errorf("Name this(%v) Not Equal that(%v)", this.Name, that1.Name)
	}
	if this.Age != that1.Age {
		return fmt.Errorf("Age this(%v) Not Equal that(%v)", this.Age, that1.Age)
	}
	return nil
}
func (this *Kept) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Kept)
	if !ok {
		that2, ok := that.(Kept)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Age != that1.Age {
		return false
	}
	return true
}
func (m *Dropped) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Dropped) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Dropped) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Age != 0 {
		i = encodeVarintTypedeclall(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypedeclall(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DroppedWithoutGetters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DroppedWithoutGetters) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DroppedWithoutGetters) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Width != 0 {
		i = encodeVarintTypedeclall(dAtA, i, uint64(m.Width))
		i--
		dAtA[i] = 0x10
	}
	if m.Height != 0 {
		i = encodeVarintTypedeclall(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Kept) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Kept) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Kept) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Age != 0 {
		i = encodeVarintTypedeclall(dAtA, i, uint64(m.Age))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypedeclall(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypedeclall(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypedeclall(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedDropped(r randyTypedeclall, easy bool) *Dropped {
	this := &Dropped{}
	this.Name = string(randStringTypedeclall(r))
	this.Age = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Age *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedDroppedWithoutGetters(r randyTypedeclall, easy bool) *DroppedWithoutGetters {
	this := &DroppedWithoutGetters{}
	this.Height = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Height *= -1
	}
	this.Width = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Width *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

func NewPopulatedKept(r randyTypedeclall, easy bool) *Kept {
	this := &Kept{}
	this.Name = string(randStringTypedeclall(r))
	this.Age = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Age *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyTypedeclall interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTypedeclall(r randyTypedeclall) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTypedeclall(r randyTypedeclall) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTypedeclall(r)
	}
	return string(tmps)
}
func randUnrecognizedTypedeclall(r randyTypedeclall, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTypedeclall(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTypedeclall(dAtA []byte, r randyTypedeclall, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTypedeclall(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTypedeclall(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTypedeclall(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTypedeclall(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTypedeclall(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTypedeclall(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTypedeclall(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Dropped) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypedeclall(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovTypedeclall(uint64(m.Age))
	}
	return n
}

func (m *DroppedWithoutGetters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovTypedeclall(uint64(m.Height))
	}
	if m.Width != 0 {
		n += 1 + sovTypedeclall(uint64(m.Width))
	}
	return n
}

func (m *Kept) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypedeclall(uint64(l))
	}
	if m.Age != 0 {
		n += 1 + sovTypedeclall(uint64(m.Age))
	}
	return n
}

func sovTypedeclall(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypedeclall(x uint64) (n int) {
	return sovTypedeclall(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Dropped) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypedeclall
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
			return fmt.Errorf("proto: Dropped: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dropped: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypedeclall
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
				return ErrInvalidLengthTypedeclall
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypedeclall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypedeclall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypedeclall(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypedeclall
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypedeclall
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
func (m *DroppedWithoutGetters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypedeclall
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
			return fmt.Errorf("proto: DroppedWithoutGetters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DroppedWithoutGetters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypedeclall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypedeclall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Width |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypedeclall(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypedeclall
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypedeclall
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
func (m *Kept) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypedeclall
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
			return fmt.Errorf("proto: Kept: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Kept: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypedeclall
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
				return ErrInvalidLengthTypedeclall
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypedeclall
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Age", wireType)
			}
			m.Age = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypedeclall
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Age |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypedeclall(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypedeclall
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypedeclall
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
func skipTypedeclall(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypedeclall
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
					return 0, ErrIntOverflowTypedeclall
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
					return 0, ErrIntOverflowTypedeclall
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
				return 0, ErrInvalidLengthTypedeclall
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypedeclall
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypedeclall
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypedeclall        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypedeclall          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypedeclall = fmt.Errorf("proto: unexpected end of group")
)
