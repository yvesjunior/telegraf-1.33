// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/sensu/sensu-go/api/core/v2/tessen.proto

package v2

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_golang_protobuf_proto "github.com/golang/protobuf/proto"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// TessenConfig is the representation of a tessen configuration.
type TessenConfig struct {
	// OptOut is the opt-out status of the tessen configuration
	OptOut               bool     `protobuf:"varint,1,opt,name=opt_out,json=optOut,proto3" json:"opt_out"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TessenConfig) Reset()         { *m = TessenConfig{} }
func (m *TessenConfig) String() string { return proto.CompactTextString(m) }
func (*TessenConfig) ProtoMessage()    {}
func (*TessenConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_639d7e72499ed3e0, []int{0}
}
func (m *TessenConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TessenConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TessenConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TessenConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TessenConfig.Merge(m, src)
}
func (m *TessenConfig) XXX_Size() int {
	return m.Size()
}
func (m *TessenConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TessenConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TessenConfig proto.InternalMessageInfo

func init() {
	proto.RegisterType((*TessenConfig)(nil), "sensu.core.v2.TessenConfig")
}

func init() {
	proto.RegisterFile("github.com/sensu/sensu-go/api/core/v2/tessen.proto", fileDescriptor_639d7e72499ed3e0)
}

var fileDescriptor_639d7e72499ed3e0 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0x4e, 0xcd, 0x2b, 0x2e, 0x85, 0x90, 0xba, 0xe9,
	0xf9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0xc9, 0xf9, 0x45, 0xa9, 0xfa, 0x65, 0x46, 0xfa, 0x25, 0xa9,
	0xc5, 0xc5, 0xa9, 0x79, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x60, 0x25, 0x7a, 0x20,
	0x39, 0xbd, 0x32, 0x23, 0x29, 0x13, 0x24, 0x23, 0xd2, 0xf3, 0xd3, 0xf3, 0xf5, 0xc1, 0xaa, 0x92,
	0x4a, 0xd3, 0x1c, 0xca, 0x0c, 0xf5, 0x8c, 0xf5, 0x0c, 0xc1, 0x82, 0x60, 0x31, 0x30, 0x0b, 0x62,
	0x88, 0x92, 0x1d, 0x17, 0x4f, 0x08, 0xd8, 0x50, 0xe7, 0xfc, 0xbc, 0xb4, 0xcc, 0x74, 0x21, 0x15,
	0x2e, 0xf6, 0xfc, 0x82, 0x92, 0xf8, 0xfc, 0xd2, 0x12, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x0e, 0x27,
	0xee, 0x57, 0xf7, 0xe4, 0x61, 0x42, 0x41, 0x6c, 0xf9, 0x05, 0x25, 0xfe, 0xa5, 0x25, 0x56, 0x1c,
	0x1d, 0x0b, 0xe4, 0x19, 0x56, 0x2c, 0x90, 0x67, 0x74, 0x52, 0xf8, 0xf1, 0x50, 0x8e, 0x71, 0xc5,
	0x23, 0x39, 0xc6, 0x1d, 0x8f, 0xe4, 0x18, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1,
	0xc1, 0x23, 0x39, 0xc6, 0x19, 0x8f, 0xe5, 0x18, 0xa2, 0x98, 0xca, 0x8c, 0x92, 0xd8, 0xc0, 0x16,
	0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe2, 0x0e, 0x9b, 0x27, 0xe3, 0x00, 0x00, 0x00,
}

func (this *TessenConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TessenConfig)
	if !ok {
		that2, ok := that.(TessenConfig)
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
	if this.OptOut != that1.OptOut {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

type TessenConfigFace interface {
	Proto() github_com_golang_protobuf_proto.Message
	GetOptOut() bool
}

func (this *TessenConfig) Proto() github_com_golang_protobuf_proto.Message {
	return this
}

func (this *TessenConfig) TestProto() github_com_golang_protobuf_proto.Message {
	return NewTessenConfigFromFace(this)
}

func (this *TessenConfig) GetOptOut() bool {
	return this.OptOut
}

func NewTessenConfigFromFace(that TessenConfigFace) *TessenConfig {
	this := &TessenConfig{}
	this.OptOut = that.GetOptOut()
	return this
}

func (m *TessenConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TessenConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TessenConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.OptOut {
		i--
		if m.OptOut {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTessen(dAtA []byte, offset int, v uint64) int {
	offset -= sovTessen(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedTessenConfig(r randyTessen, easy bool) *TessenConfig {
	this := &TessenConfig{}
	this.OptOut = bool(bool(r.Intn(2) == 0))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedTessen(r, 2)
	}
	return this
}

type randyTessen interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneTessen(r randyTessen) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringTessen(r randyTessen) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneTessen(r)
	}
	return string(tmps)
}
func randUnrecognizedTessen(r randyTessen, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldTessen(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldTessen(dAtA []byte, r randyTessen, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateTessen(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateTessen(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateTessen(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateTessen(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateTessen(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateTessen(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateTessen(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *TessenConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.OptOut {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTessen(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTessen(x uint64) (n int) {
	return sovTessen(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TessenConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTessen
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
			return fmt.Errorf("proto: TessenConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TessenConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OptOut", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTessen
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.OptOut = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTessen(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTessen
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTessen(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTessen
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
					return 0, ErrIntOverflowTessen
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
					return 0, ErrIntOverflowTessen
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
				return 0, ErrInvalidLengthTessen
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTessen
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTessen
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTessen        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTessen          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTessen = fmt.Errorf("proto: unexpected end of group")
)