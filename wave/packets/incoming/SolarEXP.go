package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SolarEXPPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 int64
	Variable2 int64
}

func (packet *SolarEXPPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *SolarEXPPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable1, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable2, b.Index())
}
