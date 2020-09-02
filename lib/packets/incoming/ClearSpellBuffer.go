package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ClearSpellBufferPacket struct {
	*wave.DefaultPacket
	SpellNum byte
}

func (packet *ClearSpellBufferPacket) Read(b buffer.PacketBuffer) {
	packet.SpellNum = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ClearSpellBufferPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.SpellNum, b.Index())
}
