package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SendTimePacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 bool
}

func (packet *SendTimePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *SendTimePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable1, b.Index())
}
