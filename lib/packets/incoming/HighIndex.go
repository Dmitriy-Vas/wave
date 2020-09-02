package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type HighIndexPacket struct {
	*wave.DefaultPacket
	HighIndex int64
}

func (packet *HighIndexPacket) Read(b buffer.PacketBuffer) {
	packet.HighIndex = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *HighIndexPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.HighIndex, b.Index())
}
