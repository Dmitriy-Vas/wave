package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ProfessionStartPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 byte
	Variable2 int32
	Variable3 int32
	Variable4 int32
}

func (packet *ProfessionStartPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *ProfessionStartPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable4, b.Index())
}
