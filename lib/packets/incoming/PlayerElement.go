package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type PlayerElementPacket struct {
	*wave.DefaultPacket
	Variable1    int32
	ElementCount int32
	Elements     []lib.PlayerElementRec
}

func (packet *PlayerElementPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.ElementCount = b.ReadInt(b.Bytes(), b.Index())
	if packet.ElementCount > 0 {
		packet.Elements = make([]lib.PlayerElementRec, packet.ElementCount)
		for i, _ := range packet.Elements {
			packet.Elements[i] = lib.PlayerElementRec{
				Num:  b.ReadInt(b.Bytes(), b.Index()),
				Icon: b.ReadInt(b.Bytes(), b.Index()),
			}
		}
	}
}

func (packet *PlayerElementPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.ElementCount, b.Index())
	if packet.ElementCount > 0 {
		for _, element := range packet.Elements {
			b.WriteInt(b.Bytes(), element.Num, b.Index())
			b.WriteInt(b.Bytes(), element.Icon, b.Index())
		}
	}
}
