package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateResourcePacket struct {
	*wave.DefaultPacket
	Variable0 int32
}

func (packet *UpdateResourcePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadInt(b.Bytes(), b.Index())
	// TODO resource data
}

func (packet *UpdateResourcePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable0, b.Index())
	// TODO resource data
}
