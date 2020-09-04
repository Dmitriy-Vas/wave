package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SpawnItemPacket struct {
	*wave.DefaultPacket
	TmpItem   int32
	TmpAmount int32
}

func (packet *SpawnItemPacket) Read(b buffer.PacketBuffer) {
	packet.TmpItem = b.ReadInt(b.Bytes(), b.Index())
	packet.TmpAmount = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *SpawnItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.TmpItem, b.Index())
	b.WriteInt(b.Bytes(), packet.TmpAmount, b.Index())
}
