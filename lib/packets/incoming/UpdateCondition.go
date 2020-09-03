package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateConditionPacket struct {
	*wave.DefaultPacket
	Variable0 int32
}

func (packet *UpdateConditionPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadInt(b.Bytes(), b.Index())
	if v := packet.Variable0; v != 0 || v <= 400 {
		// TODO condition data
	}
}

func (packet *UpdateConditionPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable0, b.Index())
	if v := packet.Variable0; v != 0 || v <= 400 {
		// TODO condition data
	}
}
