package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerXYMapPacket struct {
	*wave.DefaultPacket
	Index int32
	X     int32
	Y     int32
	Dir   byte
}

func (packet *PlayerXYMapPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.X = b.ReadInt(b.Bytes(), b.Index())
	packet.Y = b.ReadInt(b.Bytes(), b.Index())
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerXYMapPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteInt(b.Bytes(), packet.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Y, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
}
