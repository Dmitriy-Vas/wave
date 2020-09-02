package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type MapItemDataPacket struct {
	*wave.DefaultPacket
	ItemNum   int32
	PlayerNum int32
}

func (packet *MapItemDataPacket) Read(b buffer.PacketBuffer) {
	packet.ItemNum = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *MapItemDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ItemNum, b.Index())
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
}
