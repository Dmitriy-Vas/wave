package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type HotbarChangePacket struct {
	*wave.DefaultPacket
	Type      byte
	Slot      int32
	HotbarNum int32
}

func (packet *HotbarChangePacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Slot = b.ReadInt(b.Bytes(), b.Index())
	packet.HotbarNum = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *HotbarChangePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteInt(b.Bytes(), packet.Slot, b.Index())
	b.WriteInt(b.Bytes(), packet.HotbarNum, b.Index())
}
