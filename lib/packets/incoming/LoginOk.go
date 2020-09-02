package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type LoginOkPacket struct {
	*wave.DefaultPacket
	MyIndex int32
	Class   int32
}

func (packet *LoginOkPacket) Read(b buffer.PacketBuffer) {
	packet.MyIndex = b.ReadInt(b.Bytes(), b.Index())
	packet.Class = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *LoginOkPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MyIndex, b.Index())
	b.WriteInt(b.Bytes(), packet.Class, b.Index())
}
