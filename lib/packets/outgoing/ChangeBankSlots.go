package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ChangeBankSlotsPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ChangeBankSlotsPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ChangeBankSlotsPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ChangeBankSlotsPacket) SetSend(value bool) {
	d.Send = value
}

type ChangeBankSlotsPacket struct {
	ID      int64
	Send    bool
	OldSlot byte
	NewSlot byte
}

func (packet *ChangeBankSlotsPacket) Read(b buffer.PacketBuffer) {
	packet.OldSlot = b.ReadByte(b.Bytes(), b.Index())
	packet.NewSlot = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *ChangeBankSlotsPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.OldSlot, b.Index())
	b.WriteByte(b.Bytes(), packet.NewSlot, b.Index())
}
