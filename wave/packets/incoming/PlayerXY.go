package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerXYPacket struct {
	*wave.DefaultPacket
	X   int32
	Y   int32
	Dir byte
}

func (p *PlayerXYPacket) Read(b buffer.PacketBuffer) {
	p.X = b.ReadInt(b.Bytes(), b.Index())
	p.Y = b.ReadInt(b.Bytes(), b.Index())
	p.Dir = b.ReadByte(b.Bytes(), b.Index())
}

func (p *PlayerXYPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), p.X, b.Index())
	b.WriteInt(b.Bytes(), p.Y, b.Index())
	b.WriteByte(b.Bytes(), p.Dir, b.Index())
}
