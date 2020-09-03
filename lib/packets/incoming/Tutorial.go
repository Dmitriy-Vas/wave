package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type TutorialPacket struct {
	*wave.DefaultPacket
	TutorialNum int32
}

func (packet *TutorialPacket) Read(b buffer.PacketBuffer) {
	packet.TutorialNum = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *TutorialPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.TutorialNum, b.Index())
}
