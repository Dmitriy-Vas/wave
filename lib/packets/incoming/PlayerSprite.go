package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerSpritePacket struct {
	*wave.DefaultPacket
	PlayerNum int32
	Sprite    int32
	Polymorph int32
}

func (packet *PlayerSpritePacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Sprite = b.ReadInt(b.Bytes(), b.Index())
	packet.Polymorph = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerSpritePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Sprite, b.Index())
	b.WriteInt(b.Bytes(), packet.Polymorph, b.Index())
}
