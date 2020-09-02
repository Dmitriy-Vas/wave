package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerDeathPacket struct {
	*wave.DefaultPacket
	PlayerNum int32
	Ghost     byte
}

func (packet *PlayerDeathPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Ghost = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerDeathPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteByte(b.Bytes(), packet.Ghost, b.Index())
}
