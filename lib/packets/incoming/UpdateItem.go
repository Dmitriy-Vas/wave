package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateItemPacket struct {
	*wave.DefaultPacket
	ItemNum int32
}

func (packet *UpdateItemPacket) Read(b buffer.PacketBuffer) {
	packet.ItemNum = b.ReadInt(b.Bytes(), b.Index())
	// TODO item data
}

func (packet *UpdateItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ItemNum, b.Index())
	// TODO item data
}
