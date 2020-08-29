package buffer

import (
	"encoding/binary"
	"math"
)

type PacketReader interface {
	Init(args []interface{})
	ReadBool(data []byte, index uint32) bool                    // 1 byte
	ReadByte(data []byte, index uint32) byte                    // 1 byte
	ReadFloat(data []byte, index uint32) float32                // 4 bytes
	ReadDouble(data []byte, index uint32) float64               // 8 bytes
	ReadInt(data []byte, index uint32) int32                    // 4 bytes
	ReadUint(data []byte, index uint32) uint32                  // 4 bytes
	ReadLong(data []byte, index uint32) int64                   // 8 bytes
	ReadULong(data []byte, index uint32) uint64                 // 8 bytes
	ReadShort(data []byte, index uint32) int16                  // 2 byte
	ReadUShort(data []byte, index uint32) uint16                // 2 byte
	ReadVector2(data []byte, index uint32) Vector2              // 8 bytes
	ReadString(data []byte, index uint32, length uint32) string // specific length
	ReadBytes(data []byte, index uint32, length uint32) []byte  // specific length
}

type DefaultReader struct {
	order binary.ByteOrder
}

func (dr *DefaultReader) Init(args []interface{}) {
	var order binary.ByteOrder = binary.LittleEndian
	if args != nil {
		order = args[0].(binary.ByteOrder)
	}
	dr.order = order
}

func (dr *DefaultReader) ReadBool(data []byte, index uint32) bool {
	return dr.ReadByte(data, index) != 0
}

func (dr *DefaultReader) ReadByte(data []byte, index uint32) byte {
	return data[index : index+1][0]
}

func (dr *DefaultReader) ReadFloat(data []byte, index uint32) float32 {
	return math.Float32frombits(dr.ReadUint(data, index))
}

func (dr *DefaultReader) ReadDouble(data []byte, index uint32) float64 {
	return math.Float64frombits(dr.ReadULong(data, index))
}

func (dr *DefaultReader) ReadInt(data []byte, index uint32) int32 {
	return int32(dr.ReadUint(data, index))
}

func (dr *DefaultReader) ReadUint(data []byte, index uint32) uint32 {
	return dr.order.Uint32(data[index:])
}

func (dr *DefaultReader) ReadLong(data []byte, index uint32) int64 {
	return int64(dr.ReadULong(data, index))
}

func (dr *DefaultReader) ReadULong(data []byte, index uint32) uint64 {
	return dr.order.Uint64(data[index:])
}

func (dr *DefaultReader) ReadShort(data []byte, index uint32) int16 {
	return int16(dr.ReadUShort(data, index))
}

func (dr *DefaultReader) ReadUShort(data []byte, index uint32) uint16 {
	return dr.order.Uint16(data[index:])
}

func (dr *DefaultReader) ReadVector2(data []byte, index uint32) Vector2 {
	return Vector2{
		First:  dr.ReadInt(data, index),
		Second: dr.ReadInt(data, index+4),
	}
}

func (dr *DefaultReader) ReadString(data []byte, index uint32, length uint32) string {
	return string(dr.ReadBytes(data, index, length))
}

func (dr *DefaultReader) ReadBytes(data []byte, index uint32, length uint32) []byte {
	return data[index : index+length]
}
