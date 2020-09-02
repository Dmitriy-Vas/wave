package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type TravelInfoPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 bool
}

func (packet *TravelInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *TravelInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable1, b.Index())
}
