package wave

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"reflect"
)

const (
	Size64Bits = 8
	Size32Bits = 4
	Size16Bits = 2
	Size8Bits  = 1
)

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

// TODO still can't exclude reflect from all aspects, need to find better way to work with packets
// Creates new structure instance from its type
func InitializeStruct(t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.Ptr:
		return reflect.New(t.Elem()).Interface()
	case reflect.Struct:
		return reflect.New(t).Interface()
	}
	return nil
}
