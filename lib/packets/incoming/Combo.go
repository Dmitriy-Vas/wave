package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ComboPacket struct {
	*wave.DefaultPacket
	Variable0 int64
}

func (packet *ComboPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *ComboPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
}
