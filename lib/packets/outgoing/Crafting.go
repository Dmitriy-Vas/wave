package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type CraftingPacket struct {
	*wave.DefaultPacket
	CraftingNum int32
}

func (packet *CraftingPacket) Read(b buffer.PacketBuffer) {
	packet.CraftingNum = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *CraftingPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.CraftingNum, b.Index())
}
