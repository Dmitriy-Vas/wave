package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SpellsPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 bool
	Variable2 int32
	Variable3 int32
	Variable4 int64
	Variable5 bool
}

func (packet *SpellsPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *SpellsPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable4, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable5, b.Index())
}
