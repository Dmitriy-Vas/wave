package buffer

import (
	"encoding/binary"
	"math"
)

type PacketWriter interface {
	Init(args []interface{})
	WriteBool(data []byte, value bool, index uint32)       // 1 byte
	WriteByte(data []byte, value byte, index uint32)       // 1 byte
	WriteFloat(data []byte, value float32, index uint32)   // 4 bytes
	WriteDouble(data []byte, value float64, index uint32)  // 8 bytes
	WriteInt(data []byte, value int32, index uint32)       // 4 bytes
	WriteUint(data []byte, value uint32, index uint32)     // 4 bytes
	WriteLong(data []byte, value int64, index uint32)      // 8 bytes
	WriteULong(data []byte, value uint64, index uint32)    // 8 bytes
	WriteShort(data []byte, value int16, index uint32)     // 2 byte
	WriteUShort(data []byte, value uint16, index uint32)   // 2 byte
	WriteVector2(data []byte, value Vector2, index uint32) // 8 bytes
	WriteString(data []byte, value string, index uint32)   // specific length
	WriteBytes(data []byte, value []byte, index uint32)    // specific length
}

type DefaultWriter struct {
	order binary.ByteOrder
}

func (dw *DefaultWriter) Init(args []interface{}) {
	var order binary.ByteOrder = binary.LittleEndian
	if args != nil {
		order = args[0].(binary.ByteOrder)
	}
	dw.order = order
}

func (dw *DefaultWriter) WriteBool(data []byte, value bool, index uint32) {
	var _value byte = 0
	if value {
		_value = 1
	}
	dw.WriteByte(data, _value, index)
}

func (dw *DefaultWriter) WriteByte(data []byte, value byte, index uint32) {
	data[index] = value
}

func (dw *DefaultWriter) WriteFloat(data []byte, value float32, index uint32) {
	dw.WriteUint(data, math.Float32bits(value), index)
}

func (dw *DefaultWriter) WriteDouble(data []byte, value float64, index uint32) {
	dw.WriteULong(data, math.Float64bits(value), index)
}

func (dw *DefaultWriter) WriteInt(data []byte, value int32, index uint32) {
	dw.WriteUint(data, uint32(value), index)
}

func (dw *DefaultWriter) WriteUint(data []byte, value uint32, index uint32) {
	dw.order.PutUint32(data[index:], value)
}

func (dw *DefaultWriter) WriteLong(data []byte, value int64, index uint32) {
	dw.WriteULong(data, uint64(value), index)
}

func (dw *DefaultWriter) WriteULong(data []byte, value uint64, index uint32) {
	dw.order.PutUint64(data[index:], value)
}

func (dw *DefaultWriter) WriteShort(data []byte, value int16, index uint32) {
	dw.WriteUShort(data, uint16(value), index)
}

func (dw *DefaultWriter) WriteUShort(data []byte, value uint16, index uint32) {
	dw.order.PutUint16(data[index:], value)
}

func (dw *DefaultWriter) WriteVector2(data []byte, value Vector2, index uint32) {
	dw.WriteInt(data, value.First, index)
	dw.WriteInt(data, value.Second, index+4)
}

func (dw *DefaultWriter) WriteString(data []byte, value string, index uint32) {
	dw.WriteBytes(data, []byte(value), index)
}

func (dw *DefaultWriter) WriteBytes(data []byte, value []byte, index uint32) {
	for i, b := range value {
		data[index+uint32(i)] = b
	}
}
