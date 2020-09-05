package wave

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"reflect"
)

const (
	// Size64Bits specifies size for long number (int64 and uint64)
	Size64Bits = 8
	// Size32Bits specifies size for int number (int32 and uint32)
	Size32Bits = 4
	// Size16Bits specifies size for short number (int16 and uint16)
	Size16Bits = 2
	// Size8Bits specifies size for byte and bool (int8 and uint8)
	Size8Bits = 1
)

// ReadNumber reads any number with any type, based on specified size (in bytes).
// For example, if size=32bits, int (4 bytes) will be read from the buffer.
func ReadNumber(buffer buffer.PacketBuffer, size uint64) (result int64) {
	switch size {
	case Size64Bits:
		result = buffer.ReadLong(buffer.Bytes(), buffer.Index())
	case Size32Bits:
		result = int64(buffer.ReadInt(buffer.Bytes(), buffer.Index()))
	case Size16Bits:
		result = int64(buffer.ReadShort(buffer.Bytes(), buffer.Index()))
	case Size8Bits:
		result = int64(buffer.ReadByte(buffer.Bytes(), buffer.Index()))
	}
	return result
}

// WriteNumber writes any number with any type, based on specified size (in bytes).
// For example, if size=32bits, int (4 bytes) will be written to the buffer.
func WriteNumber(buffer buffer.PacketBuffer, size uint64, value int64) {
	switch size {
	case Size64Bits:
		buffer.WriteLong(buffer.Bytes(), value, buffer.Index())
	case Size32Bits:
		buffer.WriteInt(buffer.Bytes(), int32(value), buffer.Index())
	case Size16Bits:
		buffer.WriteShort(buffer.Bytes(), int16(value), buffer.Index())
	case Size8Bits:
		buffer.WriteByte(buffer.Bytes(), byte(value), buffer.Index())
	}
}

// InitializeStruct Creates new instance from type.
func InitializeStruct(t reflect.Type) interface{} {
	// TODO still can't exclude reflect from all aspects, need to find better way to work with packets
	switch t.Kind() {
	case reflect.Ptr:
		return reflect.New(t.Elem()).Interface()
	case reflect.Struct:
		return reflect.New(t).Interface()
	}
	return nil
}
