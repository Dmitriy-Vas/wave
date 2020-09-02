package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ArenaBonusPacket struct {
	*wave.DefaultPacket
	Variable1 bool
	Variable2 byte
}

func (packet *ArenaBonusPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ArenaBonusPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
}
