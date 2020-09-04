package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type TradeItemPacket struct {
	*wave.DefaultPacket
	InvSlot byte
	Amount  int64
	IsSkull bool
}

func (packet *TradeItemPacket) Read(b buffer.PacketBuffer) {
	packet.InvSlot = b.ReadByte(b.Bytes(), b.Index())
	packet.Amount = b.ReadLong(b.Bytes(), b.Index())
	packet.IsSkull = b.ReadBool(b.Bytes(), b.Index())

}

func (packet *TradeItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvSlot, b.Index())
	b.WriteLong(b.Bytes(), packet.Amount, b.Index())
	b.WriteBool(b.Bytes(), packet.IsSkull, b.Index())
}
