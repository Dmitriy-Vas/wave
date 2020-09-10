package buffer

import (
	"encoding/binary"
	"math"
)

type PacketWriter interface {
	WriteBool(data []byte, value bool, index uint64)      // 1 byte
	WriteByte(data []byte, value byte, index uint64)      // 1 byte
	WriteFloat(data []byte, value float32, index uint64)  // 4 bytes
	WriteDouble(data []byte, value float64, index uint64) // 8 bytes
	WriteInt(data []byte, value int32, index uint64)      // 4 bytes
	WriteUint(data []byte, value uint32, index uint64)    // 4 bytes
	WriteLong(data []byte, value int64, index uint64)     // 8 bytes
	WriteULong(data []byte, value uint64, index uint64)   // 8 bytes
	WriteShort(data []byte, value int16, index uint64)    // 2 bytes
	WriteUShort(data []byte, value uint16, index uint64)  // 2 bytes
	WriteString(data []byte, value string, index uint64)  // specific length
	WriteBytes(data []byte, value []byte, index uint64)   // specific length
}

type DefaultWriter struct {
	Order binary.ByteOrder
}

func (dw *DefaultWriter) WriteBool(data []byte, value bool, index uint64) {
	var _value byte = 0
	if value {
		_value = 1
	}
	dw.WriteByte(data, _value, index)
}

func (dw *DefaultWriter) WriteByte(data []byte, value byte, index uint64) {
	data[index] = value
}

func (dw *DefaultWriter) WriteFloat(data []byte, value float32, index uint64) {
	dw.WriteUint(data, math.Float32bits(value), index)
}

func (dw *DefaultWriter) WriteDouble(data []byte, value float64, index uint64) {
	dw.WriteULong(data, math.Float64bits(value), index)
}

func (dw *DefaultWriter) WriteInt(data []byte, value int32, index uint64) {
	dw.WriteUint(data, uint32(value), index)
}

func (dw *DefaultWriter) WriteUint(data []byte, value uint32, index uint64) {
	dw.Order.PutUint32(data[index:], value)
}

func (dw *DefaultWriter) WriteLong(data []byte, value int64, index uint64) {
	dw.WriteULong(data, uint64(value), index)
}

func (dw *DefaultWriter) WriteULong(data []byte, value uint64, index uint64) {
	dw.Order.PutUint64(data[index:], value)
}

func (dw *DefaultWriter) WriteShort(data []byte, value int16, index uint64) {
	dw.WriteUShort(data, uint16(value), index)
}

func (dw *DefaultWriter) WriteUShort(data []byte, value uint16, index uint64) {
	dw.Order.PutUint16(data[index:], value)
}

func (dw *DefaultWriter) WriteString(data []byte, value string, index uint64) {
	dw.WriteBytes(data, []byte(value), index)
}

func (dw *DefaultWriter) WriteBytes(data []byte, value []byte, index uint64) {
	copy(data[index:], value)
}
