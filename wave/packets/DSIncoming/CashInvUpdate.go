package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type CashInvUpdatePacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 byte
	Variable2 buffer.Vector2
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &CashInvUpdatePacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *CashInvUpdatePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadVector2(b.Bytes(), b.Index())
}

func (packet *CashInvUpdatePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteVector2(b.Bytes(), packet.Variable2, b.Index())
}
