package buffer

import (
	"encoding/binary"
	"math"
)

type PacketReader interface {
	ReadBool(data []byte, index uint64) bool                    // 1 byte
	ReadByte(data []byte, index uint64) byte                    // 1 byte
	ReadFloat(data []byte, index uint64) float32                // 4 bytes
	ReadDouble(data []byte, index uint64) float64               // 8 bytes
	ReadInt(data []byte, index uint64) int32                    // 4 bytes
	ReadUint(data []byte, index uint64) uint32                  // 4 bytes
	ReadLong(data []byte, index uint64) int64                   // 8 bytes
	ReadULong(data []byte, index uint64) uint64                 // 8 bytes
	ReadShort(data []byte, index uint64) int16                  // 2 bytes
	ReadUShort(data []byte, index uint64) uint16                // 2 bytes
	ReadString(data []byte, index uint64, length uint64) string // specific length
	ReadBytes(data []byte, index uint64, length uint64) []byte  // specific length
}

type DefaultReader struct {
	Order binary.ByteOrder
}

func (dr *DefaultReader) ReadBool(data []byte, index uint64) bool {
	return dr.ReadByte(data, index) != 0
}

func (dr *DefaultReader) ReadByte(data []byte, index uint64) byte {
	return data[index]
}

func (dr *DefaultReader) ReadFloat(data []byte, index uint64) float32 {
	return math.Float32frombits(dr.ReadUint(data, index))
}

func (dr *DefaultReader) ReadDouble(data []byte, index uint64) float64 {
	return math.Float64frombits(dr.ReadULong(data, index))
}

func (dr *DefaultReader) ReadInt(data []byte, index uint64) int32 {
	return int32(dr.ReadUint(data, index))
}

func (dr *DefaultReader) ReadUint(data []byte, index uint64) uint32 {
	return dr.Order.Uint32(data[index:])
}

func (dr *DefaultReader) ReadLong(data []byte, index uint64) int64 {
	return int64(dr.ReadULong(data, index))
}

func (dr *DefaultReader) ReadULong(data []byte, index uint64) uint64 {
	return dr.Order.Uint64(data[index:])
}

func (dr *DefaultReader) ReadShort(data []byte, index uint64) int16 {
	return int16(dr.ReadUShort(data, index))
}

func (dr *DefaultReader) ReadUShort(data []byte, index uint64) uint16 {
	return dr.Order.Uint16(data[index:])
}

func (dr *DefaultReader) ReadString(data []byte, index uint64, length uint64) string {
	return string(dr.ReadBytes(data, index, length))
}

func (dr *DefaultReader) ReadBytes(data []byte, index uint64, length uint64) []byte {
	return data[index : index+length]
}
