package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerInfoPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 int32
	Variable2 int32
	Variable3 string
	Variable4 bool
}

func (packet *PlayerInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable4 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *PlayerInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteString(b.Bytes(), packet.Variable3, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable4, b.Index())
}
