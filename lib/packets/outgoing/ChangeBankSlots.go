package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ChangeBankSlotsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ChangeBankSlotsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ChangeBankSlotsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ChangeBankSlotsPacket) SetSend(value bool) {
	packet.Send = value
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
