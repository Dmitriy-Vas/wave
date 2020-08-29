package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type StunnedPacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 int32
	Variable2 bool
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &StunnedPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *StunnedPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *StunnedPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable2, b.Index())
}
