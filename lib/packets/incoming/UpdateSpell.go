package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateSpellPacket struct {
	*wave.DefaultPacket
	Variable0 int32
}

func (packet *UpdateSpellPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadInt(b.Bytes(), b.Index())
	// TODO spell data
}

func (packet *UpdateSpellPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable0, b.Index())
	// TODO spell data
}
