package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerTimerPacket struct {
	*wave.DefaultPacket
	ServerTimer  int32
	TravelActive byte
}

func (packet *PlayerTimerPacket) Read(b buffer.PacketBuffer) {
	packet.ServerTimer = b.ReadInt(b.Bytes(), b.Index())
	packet.TravelActive = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerTimerPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ServerTimer, b.Index())
	b.WriteByte(b.Bytes(), packet.TravelActive, b.Index())
}
