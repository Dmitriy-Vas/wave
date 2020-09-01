package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateSpellPacket struct {
	*wave.DefaultPacket
	Variable0 int64
}

func (packet *UpdateSpellPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *UpdateSpellPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
}
