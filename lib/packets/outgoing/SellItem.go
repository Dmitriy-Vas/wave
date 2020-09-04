package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SellItemPacket struct {
	*wave.DefaultPacket
	InvSlot  byte
	Multiple int64
}

func (packet *SellItemPacket) Read(b buffer.PacketBuffer) {
	packet.InvSlot = b.ReadByte(b.Bytes(), b.Index())
	packet.Multiple = b.ReadLong(b.Bytes(), b.Index())

}

func (packet *SellItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvSlot, b.Index())
	b.WriteLong(b.Bytes(), packet.Multiple, b.Index())
}
