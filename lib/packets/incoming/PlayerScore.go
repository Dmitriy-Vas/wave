package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerScorePacket struct {
	*wave.DefaultPacket
	SoccerScore int32
}

func (packet *PlayerScorePacket) Read(b buffer.PacketBuffer) {
	packet.SoccerScore = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerScorePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.SoccerScore, b.Index())
}
