package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerHonorPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 int32
	Variable2 int32
	Variable3
}

func (packet *PlayerHonorPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.Read(b.Bytes(), b.Index())
}

func (packet *PlayerHonorPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.Write(b.Bytes(), packet.Variable3, b.Index())
}
