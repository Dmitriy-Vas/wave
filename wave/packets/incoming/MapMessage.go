package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type MapMessagePacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1
}

func (packet *MapMessagePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.Read(b.Bytes(), b.Index())
}

func (packet *MapMessagePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.Write(b.Bytes(), packet.Variable1, b.Index())
}
