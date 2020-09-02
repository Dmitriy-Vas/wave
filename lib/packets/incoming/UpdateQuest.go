package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateQuestPacket struct {
	*wave.DefaultPacket
	Variable1 int32
}

func (packet *UpdateQuestPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	// TODO quest data
}

func (packet *UpdateQuestPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	// TODO quest data
}
