package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type NpcBuffInfoPacket struct {
	*wave.DefaultPacket
	Index int32
	Slot  byte
}

func (packet *NpcBuffInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.Slot = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *NpcBuffInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteByte(b.Bytes(), packet.Slot, b.Index())
}
