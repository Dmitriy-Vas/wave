package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"time"
)

type PlayerHonorPacket struct {
	*wave.DefaultPacket
	PlayerNum int32
	Honor     int32
	HonorDate time.Time
}

func (packet *PlayerHonorPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Honor = b.ReadInt(b.Bytes(), b.Index())
	packet.HonorDate = wrapper.ReadDate(b)
}

func (packet *PlayerHonorPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Honor, b.Index())
	wrapper.WriteDate(b, packet.HonorDate)
}
