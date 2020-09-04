package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type TitlePacket struct {
	*wave.DefaultPacket
	TitleNum int32
}

func (packet *TitlePacket) Read(b buffer.PacketBuffer) {
	packet.TitleNum = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *TitlePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.TitleNum, b.Index())
}
