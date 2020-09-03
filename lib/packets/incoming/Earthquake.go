package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type EarthquakePacket struct {
	*wave.DefaultPacket
	ShowEartquake byte
}

func (packet *EarthquakePacket) Read(b buffer.PacketBuffer) {
	packet.ShowEartquake = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *EarthquakePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.ShowEartquake, b.Index())
}
