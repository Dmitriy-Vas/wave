package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ChangeSpellSlotsPacket struct {
	*wave.DefaultPacket
	OldSlot int64
	NewSlot int64
}

func (packet *ChangeSpellSlotsPacket) Read(b buffer.PacketBuffer) {
	// TODO if modGlobals.SpellCD[checked((int)OldSlot)].CDTimer <= 0
	packet.OldSlot = b.ReadLong(b.Bytes(), b.Index())
	packet.NewSlot = b.ReadLong(b.Bytes(), b.Index())

}

func (packet *ChangeSpellSlotsPacket) Write(b buffer.PacketBuffer) {
	// TODO if modGlobals.SpellCD[checked((int)OldSlot)].CDTimer <= 0
	b.WriteLong(b.Bytes(), packet.OldSlot, b.Index())
	b.WriteLong(b.Bytes(), packet.NewSlot, b.Index())
}
