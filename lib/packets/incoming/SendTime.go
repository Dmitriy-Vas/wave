package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SendTimePacket struct {
	*wave.DefaultPacket
	DayTime bool
}

func (packet *SendTimePacket) Read(b buffer.PacketBuffer) {
	packet.DayTime = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *SendTimePacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.DayTime, b.Index())
}
