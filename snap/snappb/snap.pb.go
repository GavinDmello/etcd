// Code generated by protoc-gen-gogo.
// source: snap.proto
// DO NOT EDIT!

/*
	Package snappb is a generated protocol buffer package.

	It is generated from these files:
		snap.proto

	It has these top-level messages:
		Snapshot
*/
package snappb

import proto "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"
import json "encoding/json"
import math "math"

// discarding unused import gogoproto "code.google.com/p/gogoprotobuf/gogoproto/gogo.pb"

import io "io"
import code_google_com_p_gogoprotobuf_proto "github.com/coreos/etcd/third_party/code.google.com/p/gogoprotobuf/proto"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type Snapshot struct {
	Crc              uint32 `protobuf:"varint,1,req,name=crc" json:"crc"`
	Data             []byte `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Snapshot) Reset()         { *m = Snapshot{} }
func (m *Snapshot) String() string { return proto.CompactTextString(m) }
func (*Snapshot) ProtoMessage()    {}

func init() {
}
func (m *Snapshot) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				m.Crc |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return code_google_com_p_gogoprotobuf_proto.ErrWrongType
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, data[index:postIndex]...)
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := code_google_com_p_gogoprotobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Snapshot) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovSnap(uint64(m.Crc))
	if m.Data != nil {
		l = len(m.Data)
		n += 1 + l + sovSnap(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSnap(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSnap(x uint64) (n int) {
	return sovSnap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Snapshot) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Snapshot) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintSnap(data, i, uint64(m.Crc))
	if m.Data != nil {
		data[i] = 0x12
		i++
		i = encodeVarintSnap(data, i, uint64(len(m.Data)))
		i += copy(data[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}
func encodeFixed64Snap(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Snap(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSnap(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}