package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type EnhancementPacket struct {
	*wave.DefaultPacket
	InvSlot    byte
	TargetSlot byte
}

func (packet *EnhancementPacket) Read(b buffer.PacketBuffer) {
	packet.InvSlot = b.ReadByte(b.Bytes(), b.Index())
	packet.TargetSlot = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *EnhancementPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvSlot, b.Index())
	b.WriteByte(b.Bytes(), packet.TargetSlot, b.Index())
}
