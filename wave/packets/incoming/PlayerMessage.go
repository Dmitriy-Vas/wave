package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerMessagePacket struct {
	*wave.DefaultPacket
	Variable0  int64
	Variable1  int32
	Variable2  int32
	Variable3  string
	Variable4  string
	Variable5  string
	Variable6  bool
	Variable7  int32
	Variable8  byte
	Variable9  byte
	Variable10 byte
}

func (packet *PlayerMessagePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable4 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable5 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable6 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable8 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable9 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable10 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerMessagePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteString(b.Bytes(), packet.Variable3, b.Index())
	b.WriteString(b.Bytes(), packet.Variable4, b.Index())
	b.WriteString(b.Bytes(), packet.Variable5, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable6, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable7, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable8, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable9, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable10, b.Index())
}
