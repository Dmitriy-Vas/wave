package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type TradeStatusPacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 string
	Variable2 byte
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &TradeStatusPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *TradeStatusPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *TradeStatusPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteString(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
}
