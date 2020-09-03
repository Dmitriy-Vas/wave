package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type TravelInfoPacket struct {
	*wave.DefaultPacket
	TravelActive bool
}

func (packet *TravelInfoPacket) Read(b buffer.PacketBuffer) {
	packet.TravelActive = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *TravelInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.TravelActive, b.Index())
}
