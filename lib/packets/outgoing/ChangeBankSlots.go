package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ChangeBankSlotsPacket struct {
	*wave.DefaultPacket
	OldSlot byte
	NewSlot byte
}

func (packet *ChangeBankSlotsPacket) Read(b buffer.PacketBuffer) {
	packet.OldSlot = b.ReadByte(b.Bytes(), b.Index())
	packet.NewSlot = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *ChangeBankSlotsPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.OldSlot, b.Index())
	b.WriteByte(b.Bytes(), packet.NewSlot, b.Index())
}
