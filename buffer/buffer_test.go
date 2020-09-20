package buffer

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

var testReader = &DefaultReader{Order: binary.LittleEndian}
var testWriter = &DefaultWriter{Order: binary.LittleEndian}

// New buffer for each test
func newTestBuffer() *DefaultBuffer {
	return &DefaultBuffer{
		PacketReader: testReader,
		PacketWriter: testWriter,
		maxLen:       255,
		len:          255,
		initLen:      8,
		buf:          make([]byte, 255),
	}
}

func TestDefaultReader_ReadBool(t *testing.T) {
	buffer := newTestBuffer()
	buffer.buf[0] = 1
	result1 := buffer.ReadBool(buffer.Bytes(), buffer.Index())
	assert.Exactly(t, true, result1)
	buffer.buf[0] = 0
	result2 := buffer.ReadBool(buffer.Bytes(), buffer.Index())
	assert.Exactly(t, false, result2)
}

func TestDefaultReader_ReadByte(t *testing.T) {
	buffer := newTestBuffer()
	buffer.buf[0] = byte(255)
	result1 := buffer.ReadByte(buffer.Bytes(), buffer.Index())
	buffer.Back(1)
	assert.Exactly(t, byte(255), result1)
	buffer.buf[0] = byte(150)
	result2 := buffer.ReadByte(buffer.Bytes(), buffer.Index())
	assert.Exactly(t, byte(150), result2)
}

func TestDefaultReader_ReadBytes(t *testing.T) {
	buffer := newTestBuffer()
	value1 := []byte{255, 255, 255, 255}
	for i, b := range value1 {
		buffer.buf[i] = b
	}
	result1 := buffer.ReadBytes(buffer.Bytes(), 0, 4)
	assert.Exactly(t, value1, result1)
}

func TestDefaultReader_ReadDouble(t *testing.T) {
	buffer := newTestBuffer()
	value1 := 3.141
	binary.LittleEndian.PutUint64(buffer.buf[:8], math.Float64bits(value1)) // [84 227 165 155 196 32 9 64]
	result1 := buffer.ReadDouble(buffer.Bytes(), 0)
	assert.Exactly(t, value1, result1)
}

func TestDefaultReader_ReadFloat(t *testing.T) {
	buffer := newTestBuffer()
	value1 := float32(3.141)
	binary.LittleEndian.PutUint32(buffer.buf[:4], math.Float32bits(value1)) // [37 6 73 64]
	result1 := buffer.ReadFloat(buffer.Bytes(), 0)
	assert.Exactly(t, value1, result1)
}

func TestDefaultReader_ReadInt(t *testing.T) {
	buffer := newTestBuffer()
	value1 := int32(999)
	binary.LittleEndian.PutUint32(buffer.buf[:4], uint32(value1))
	result1 := buffer.ReadInt(buffer.Bytes(), 0)
	assert.Exactly(t, value1, result1)
}

func TestDefaultReader_ReadLong(t *testing.T) {
	buffer := newTestBuffer()
	value1 := int64(9999)
	binary.LittleEndian.PutUint64(buffer.buf[:8], uint64(value1))
	result1 := buffer.ReadLong(buffer.Bytes(), 0)
	assert.Exactly(t, value1, result1)
}

func TestDefaultReader_ReadShort(t *testing.T) {
	buffer := newTestBuffer()
	value1 := int16(32767)
	binary.LittleEndian.PutUint16(buffer.buf[:2], uint16(value1))
	result1 := buffer.ReadShort(buffer.Bytes(), 0)
	assert.Exactly(t, value1, result1)
}

func TestDefaultReader_ReadString(t *testing.T) {
	buffer := newTestBuffer()
	value1 := "Hello World!"
	for i, b := range value1 {
		buffer.buf[i] = byte(b)
	}
	result1 := buffer.ReadString(buffer.Bytes(), 0, uint64(len(value1)))
	assert.Exactly(t, value1, result1)
}

func TestDefaultWriter_WriteBool(t *testing.T) {
	buffer := newTestBuffer()
	value1 := true
	buffer.WriteBool(buffer.buf, value1, buffer.index)
	result1 := buffer.buf[0]
	assert.Exactly(t, value1, result1 != 0)
}

func TestDefaultWriter_WriteByte(t *testing.T) {
	buffer := newTestBuffer()
	value1 := byte(190)
	buffer.WriteByte(buffer.buf, value1, buffer.index)
	result1 := buffer.buf[0]
	assert.Exactly(t, value1, result1)
}

func TestDefaultWriter_WriteBytes(t *testing.T) {
	buffer := newTestBuffer()
	value1 := []byte{190, 240, 100, 255}
	buffer.WriteBytes(buffer.buf, value1, buffer.index)
	assert.Exactly(t, value1, buffer.buf[:len(value1)])
}

func TestDefaultWriter_WriteDouble(t *testing.T) {
	buffer := newTestBuffer()
	value1 := 3.141
	buffer.WriteDouble(buffer.buf, value1, buffer.index)
	result1 := math.Float64frombits(binary.LittleEndian.Uint64(buffer.buf[:8]))
	assert.Exactly(t, value1, result1)
}

func TestDefaultWriter_WriteFloat(t *testing.T) {
	buffer := newTestBuffer()
	value1 := float32(3.141)
	buffer.WriteFloat(buffer.buf, value1, buffer.index)
	result1 := math.Float32frombits(binary.LittleEndian.Uint32(buffer.buf[:8]))
	assert.Exactly(t, value1, result1)
}

func TestDefaultWriter_WriteInt(t *testing.T) {
	buffer := newTestBuffer()
	value1 := int32(600)
	buffer.WriteInt(buffer.buf, value1, buffer.index)
	result1 := int32(binary.LittleEndian.Uint32(buffer.buf[:4]))
	assert.Exactly(t, value1, result1)
}

func TestDefaultWriter_WriteLong(t *testing.T) {
	buffer := newTestBuffer()
	value1 := int64(600)
	buffer.WriteLong(buffer.buf, value1, buffer.index)
	result1 := int64(binary.LittleEndian.Uint64(buffer.buf[:8]))
	assert.Exactly(t, value1, result1)
}

func TestDefaultWriter_WriteShort(t *testing.T) {
	buffer := newTestBuffer()
	value1 := int16(32767)
	buffer.WriteShort(buffer.buf, value1, buffer.index)
	result1 := int16(binary.LittleEndian.Uint16(buffer.buf[:4]))
	assert.Exactly(t, value1, result1)
}

func TestDefaultWriter_WriteString(t *testing.T) {
	buffer := newTestBuffer()
	value1 := "Hello World"
	buffer.WriteString(buffer.buf, value1, buffer.index)
	result1 := string(buffer.buf[:len(value1)])
	assert.Exactly(t, value1, result1)
}

func TestDefaultBuffer_Back(t *testing.T) {
	buffer := newTestBuffer()
	buffer.index = 100
	buffer.Back(50)
	assert.EqualValues(t, 50, buffer.Index())
	buffer.Back(100)
	assert.EqualValues(t, 0, buffer.Index())
}

func TestDefaultBuffer_Next(t *testing.T) {
	buffer := newTestBuffer()
	buffer.Next(100)
	assert.EqualValues(t, 100, buffer.Index())
	buffer.len = 100
	buffer.Next(50)
	assert.EqualValues(t, 150, buffer.Len())
	assert.EqualValues(t, 150, buffer.Index())
}

func TestDefaultBuffer_Bytes(t *testing.T) {
	buffer := newTestBuffer()
	bytes := buffer.Bytes()
	assert.Exactly(t, make([]byte, 255), bytes)
	assert.Len(t, bytes, 255)
}

func TestDefaultBuffer_Clone(t *testing.T) {
	buffer1 := newTestBuffer()
	buffer2 := buffer1.Clone()
	buffer2.WriteInt(buffer2.Bytes(), 400, buffer2.Index())
	result1 := buffer1.ReadInt(buffer1.Bytes(), buffer1.Index())
	assert.NotEqual(t, 400, result1)
	buffer1.WriteLong(buffer1.Bytes(), 999, buffer1.Index())
	result2 := buffer2.ReadLong(buffer2.Bytes(), buffer2.Index())
	assert.NotEqual(t, 999, result2)
	assert.Len(t, buffer1.Bytes(), len(buffer2.Bytes()))
	assert.Len(t, buffer2.Bytes(), len(buffer1.Bytes()))
}

func TestDefaultBuffer_Index(t *testing.T) {
	buffer := newTestBuffer()
	assert.Equal(t, buffer.index, buffer.Index())
	buffer.Next(100)
	assert.Equal(t, buffer.index, buffer.Index())
	buffer.Back(900)
	assert.Equal(t, buffer.index, buffer.Index())
}

func TestDefaultBuffer_Resize(t *testing.T) {
	buffer := newTestBuffer()
	assert.Panics(t, func() { buffer.Resize(1000) })
	buffer.Resize(100)
	assert.EqualValues(t, 100, buffer.Len())
}

func TestDefaultBuffer_Len(t *testing.T) {
	buffer := newTestBuffer()
	assert.EqualValues(t, 255, buffer.Len())
}

func TestDefaultBuffer_Reset(t *testing.T) {
	buffer := newTestBuffer()
	buffer.Reset()
	assert.EqualValues(t, buffer.initLen, buffer.Len())
	assert.EqualValues(t, 0, buffer.Index())
}

func TestDefaultBuffer_SetInitLength(t *testing.T) {
	buffer := newTestBuffer()
	buffer.SetInitLength(100)
	assert.EqualValues(t, 100, buffer.initLen)
}

func TestDefaultBuffer_SetMaxLength(t *testing.T) {
	buffer := newTestBuffer()
	buffer.SetMaxLength(150)
	assert.Len(t, buffer.Bytes(), 150)
	assert.EqualValues(t, buffer.maxLen, 150)
}

func TestDefaultBuffer_SetReader(t *testing.T) {
	buffer := newTestBuffer()
	assert.NotPanics(t, func() {
		buffer.SetReader(&DefaultReader{Order: binary.BigEndian})
	})
	assert.IsType(t, (*DefaultReader)(nil), buffer.PacketReader)
}

func TestDefaultBuffer_SetWriter(t *testing.T) {
	buffer := newTestBuffer()
	assert.NotPanics(t, func() {
		buffer.SetWriter(&DefaultWriter{Order: binary.BigEndian})
	})
	assert.IsType(t, (*DefaultWriter)(nil), buffer.PacketWriter)
}
