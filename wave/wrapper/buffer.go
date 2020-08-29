package wrapper

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type Buffer struct {
	buffer.PacketReader
	buffer.PacketWriter
	buffer.DefaultBuffer
}

func (b *Buffer) WriteBool(data []byte, value bool, index uint32) {
	var v int32 = 0
	if value {
		v = 1
	}
	b.AppendIfNeeded(4)
	b.Next(4)
	b.PacketWriter.WriteInt(data, v, index)
}

func (b *Buffer) WriteByte(data []byte, value byte, index uint32) {
	b.AppendIfNeeded(2)
	b.Next(2)
	b.PacketWriter.WriteShort(data, int16(value), index)
}

func (b *Buffer) WriteFloat(data []byte, value float32, index uint32) {
	b.AppendIfNeeded(4)
	b.Next(4)
	b.PacketWriter.WriteFloat(data, value, index)
}

func (b *Buffer) WriteDouble(data []byte, value float64, index uint32) {
	b.AppendIfNeeded(8)
	b.Next(8)
	b.PacketWriter.WriteDouble(data, value, index)
}

func (b *Buffer) WriteInt(data []byte, value int32, index uint32) {
	b.AppendIfNeeded(4)
	b.Next(4)
	b.PacketWriter.WriteInt(data, value, index)
}

func (b *Buffer) WriteUint(data []byte, value uint32, index uint32) {
	b.AppendIfNeeded(4)
	b.Next(4)
	b.PacketWriter.WriteUint(data, value, index)
}

func (b *Buffer) WriteLong(data []byte, value int64, index uint32) {
	b.AppendIfNeeded(8)
	b.Next(8)
	b.PacketWriter.WriteLong(data, value, index)
}

func (b *Buffer) WriteULong(data []byte, value uint64, index uint32) {
	b.AppendIfNeeded(8)
	b.Next(8)
	b.PacketWriter.WriteULong(data, value, index)
}

func (b *Buffer) WriteShort(data []byte, value int16, index uint32) {
	b.AppendIfNeeded(2)
	b.Next(2)
	b.PacketWriter.WriteShort(data, value, index)
}

func (b *Buffer) WriteUShort(data []byte, value uint16, index uint32) {
	b.AppendIfNeeded(2)
	b.Next(2)
	b.PacketWriter.WriteUShort(data, value, index)
}

func (b *Buffer) WriteVector2(data []byte, value buffer.Vector2, index uint32) {
	b.AppendIfNeeded(8)
	b.Next(8)
	b.PacketWriter.WriteVector2(data, value, index)
}

func (b *Buffer) WriteString(data []byte, value string, index uint32) {
	bytes := 4 + uint32(len(value))
	b.AppendIfNeeded(bytes)
	b.Next(bytes)
	b.PacketWriter.WriteInt(data, int32(len(value)), index)
	b.PacketWriter.WriteString(data, value, index+4)
}

func (b *Buffer) WriteBytes(data []byte, value []byte, index uint32) {
	bytes := 4 + uint32(len(value))
	b.AppendIfNeeded(bytes)
	b.Next(bytes)
	b.PacketWriter.WriteInt(data, int32(len(value)), index)
	b.PacketWriter.WriteBytes(data, value, index+4)
}

func (b *Buffer) ReadBool(data []byte, index uint32) bool {
	b.Next(4)
	return b.PacketReader.ReadInt(data, index) != 0
}

func (b *Buffer) ReadByte(data []byte, index uint32) byte {
	b.Next(2)
	return byte(b.PacketReader.ReadShort(data, index))
}

func (b *Buffer) ReadFloat(data []byte, index uint32) float32 {
	b.Next(4)
	return b.PacketReader.ReadFloat(data, index)
}

func (b *Buffer) ReadDouble(data []byte, index uint32) float64 {
	b.Next(8)
	return b.PacketReader.ReadDouble(data, index)
}

func (b *Buffer) ReadInt(data []byte, index uint32) int32 {
	b.Next(4)
	return b.PacketReader.ReadInt(data, index)
}

func (b *Buffer) ReadUint(data []byte, index uint32) uint32 {
	b.Next(4)
	return b.PacketReader.ReadUint(data, index)
}

func (b *Buffer) ReadLong(data []byte, index uint32) int64 {
	b.Next(8)
	return b.PacketReader.ReadLong(data, index)
}

func (b *Buffer) ReadULong(data []byte, index uint32) uint64 {
	b.Next(8)
	return b.PacketReader.ReadULong(data, index)
}

func (b *Buffer) ReadShort(data []byte, index uint32) int16 {
	b.Next(2)
	return b.PacketReader.ReadShort(data, index)
}

func (b *Buffer) ReadUShort(data []byte, index uint32) uint16 {
	b.Next(2)
	return b.PacketReader.ReadUShort(data, index)
}

func (b *Buffer) ReadVector2(data []byte, index uint32) buffer.Vector2 {
	b.Next(8)
	return b.PacketReader.ReadVector2(data, index)
}

func (b *Buffer) ReadString(data []byte, index uint32, length uint32) string {
	b.Next(4)
	length = uint32(b.PacketReader.ReadInt(data, index))
	b.Next(length)
	return b.PacketReader.ReadString(data, index+4, length)
}

func (b *Buffer) ReadBytes(data []byte, index uint32, length uint32) []byte {
	b.Next(4)
	length = uint32(b.PacketReader.ReadInt(data, index))
	b.Next(length)
	return b.PacketReader.ReadBytes(data, index+4, length)
}

func (b *Buffer) SetReader(reader buffer.PacketReader) {
	reader.Init(nil)
	b.PacketReader = reader
}

func (b *Buffer) SetWriter(writer buffer.PacketWriter) {
	writer.Init(nil)
	b.PacketWriter = writer
}

func (b *Buffer) Init(args []interface{}) {
	buf := buffer.DefaultBuffer{
		PacketReader: b.PacketReader,
		PacketWriter: b.PacketWriter,
	}
	buf.Init(args)
	b.DefaultBuffer = buf
}
