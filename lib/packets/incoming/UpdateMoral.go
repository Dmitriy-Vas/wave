package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateMoralPacket struct {
	*wave.DefaultPacket
	Variable1 int32
}

func (packet *UpdateMoralPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	// TODO moral data
}

func (packet *UpdateMoralPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	// TODO moral data
}
