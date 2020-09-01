package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerBuddiesPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 string
	Variable2 int32
	Variable3 int32
	Variable4 string
	Variable5 string
}

func (packet *PlayerBuddiesPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable5 = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *PlayerBuddiesPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteString(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
	b.WriteString(b.Bytes(), packet.Variable4, b.Index())
	b.WriteString(b.Bytes(), packet.Variable5, b.Index())
}
