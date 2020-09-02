package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SolarEXPPacket struct {
	*wave.DefaultPacket
	Index  int64
	Amount int64
}

func (packet *SolarEXPPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadLong(b.Bytes(), b.Index())
	packet.Amount = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *SolarEXPPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Index, b.Index())
	b.WriteLong(b.Bytes(), packet.Amount, b.Index())
}
