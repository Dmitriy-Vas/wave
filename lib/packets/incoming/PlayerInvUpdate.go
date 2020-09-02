package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerInvUpdatePacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 byte
	Variable2 int32
	Variable3 int64
	Variable4 byte
	Variable5 byte
	Variable6 bool
}

func (packet *PlayerInvUpdatePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *PlayerInvUpdatePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable3, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable4, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable5, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable6, b.Index())
}