package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type ArrowPacket struct {
	*wave.DefaultPacket
	Arrows []lib.ArrowRec
}

func (packet *ArrowPacket) Read(b buffer.PacketBuffer) {
	packet.Arrows = make([]lib.ArrowRec, 4)
	for i := range packet.Arrows {
		packet.Arrows[i] = lib.ArrowRec{
			Num:   b.ReadInt(b.Bytes(), b.Index()),
			Value: b.ReadInt(b.Bytes(), b.Index()),
		}
	}
}

func (packet *ArrowPacket) Write(b buffer.PacketBuffer) {
	for _, arrow := range packet.Arrows {
		b.WriteInt(b.Bytes(), arrow.Num, b.Index())
		b.WriteInt(b.Bytes(), arrow.Value, b.Index())
	}
}
