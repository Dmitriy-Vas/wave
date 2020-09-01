package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ArrowPacket struct {
	*wave.DefaultPacket
	Arrows [3]Arrow
}

type Arrow struct {
	Num   int32
	Value int32
}

func (packet *ArrowPacket) Read(b buffer.PacketBuffer) {
	for i := 0; i <= 3; i++ {
		packet.Arrows[i] = Arrow{
			Num:   b.ReadInt(b.Bytes(), b.Index()),
			Value: b.ReadInt(b.Bytes(), b.Index()),
		}
	}
}

func (packet *ArrowPacket) Write(b buffer.PacketBuffer) {
	for i := 0; i <= 3; i++ {
		b.WriteInt(b.Bytes(), packet.Arrows[i].Num, b.Index())
		b.WriteInt(b.Bytes(), packet.Arrows[i].Value, b.Index())
	}
}
