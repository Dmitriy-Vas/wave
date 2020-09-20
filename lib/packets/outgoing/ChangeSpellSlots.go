package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ChangeSpellSlotsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ChangeSpellSlotsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ChangeSpellSlotsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ChangeSpellSlotsPacket) SetSend(value bool) {
	packet.Send = value
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
