package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type EarthquakePacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 byte
}

func (packet *EarthquakePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *EarthquakePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
}
