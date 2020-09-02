package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerCurrentPacket struct {
	*wave.DefaultPacket
	PlayerNum     int32
	Pin           int32
	ProjectilSkin byte
	DamageSkin    byte
}

func (packet *PlayerCurrentPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Pin = b.ReadInt(b.Bytes(), b.Index())
	packet.ProjectilSkin = b.ReadByte(b.Bytes(), b.Index())
	packet.DamageSkin = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerCurrentPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Pin, b.Index())
	b.WriteByte(b.Bytes(), packet.ProjectilSkin, b.Index())
	b.WriteByte(b.Bytes(), packet.DamageSkin, b.Index())
}
