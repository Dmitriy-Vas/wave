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
		max_len:      255,
		len:          255,
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
	// buffer.buf[0] = byte(value1)
	// buffer.buf[1] = byte(value1 >> 8)
	// buffer.buf[2] = byte(value1 >> 16)
	// buffer.buf[3] = byte(value1 >> 24)
	binary.LittleEndian.PutUint32(buffer.buf[:4], uint32(value1))
	result1 := buffer.ReadInt(buffer.Bytes(), 0)
	assert.Exactly(t, value1, result1)
}

func TestDefaultReader_ReadLong(t *testing.T) {
	buffer := newTestBuffer()
	value1 := int64(9999)
	// buffer.buf[0] = byte(value1)
	// buffer.buf[1] = byte(value1 >> 8)
	// buffer.buf[2] = byte(value1 >> 16)
	// buffer.buf[3] = byte(value1 >> 24)
	// buffer.buf[4] = byte(value1 >> 32)
	// buffer.buf[5] = byte(value1 >> 40)
	// buffer.buf[6] = byte(value1 >> 48)
	// buffer.buf[7] = byte(value1 >> 56)
	binary.LittleEndian.PutUint64(buffer.buf[:8], uint64(value1))
	result1 := buffer.ReadLong(buffer.Bytes(), 0)
	assert.Exactly(t, value1, result1)
}

func TestDefaultReader_ReadShort(t *testing.T) {
	buffer := newTestBuffer()
	value1 := int16(32767)
	// buffer.buf[0] = byte(value1)
	// buffer.buf[1] = byte(value1 >> 8)
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
	value1 := float64(3.141)
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
