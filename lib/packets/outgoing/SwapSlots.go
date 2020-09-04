package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SwapSlotsPacket struct {
	*wave.DefaultPacket
	OldSlot int32
	NewSlot int32
	Type    byte
}

func (packet *SwapSlotsPacket) Read(b buffer.PacketBuffer) {
	packet.OldSlot = b.ReadInt(b.Bytes(), b.Index())
	packet.NewSlot = b.ReadInt(b.Bytes(), b.Index())
	packet.Type = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *SwapSlotsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.OldSlot, b.Index())
	b.WriteInt(b.Bytes(), packet.NewSlot, b.Index())
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
}
