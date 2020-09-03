package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type CashInvUpdatePacket struct {
	*wave.DefaultPacket
	Slot      byte
	ItemNum   int32
	ItemValue int32
}

func (packet *CashInvUpdatePacket) Read(b buffer.PacketBuffer) {
	packet.Slot = b.ReadByte(b.Bytes(), b.Index())
	b.ReadInt(b.Bytes(), b.Index())
	b.ReadInt(b.Bytes(), b.Index())
}

func (packet *CashInvUpdatePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Slot, b.Index())
	b.WriteInt(b.Bytes(), packet.ItemNum, b.Index())
	b.WriteInt(b.Bytes(), packet.ItemValue, b.Index())
}
