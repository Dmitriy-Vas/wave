package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ChangeSpellSlotsPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ChangeSpellSlotsPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ChangeSpellSlotsPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ChangeSpellSlotsPacket) SetSend(value bool) {
	d.Send = value
}

type ChangeSpellSlotsPacket struct {
	ID      int64
	Send    bool
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
