package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PartyLeavePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PartyLeavePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PartyLeavePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PartyLeavePacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type PartyLeavePacket struct {
	ID   int64
	Send bool
}

func (packet *PartyLeavePacket) Read(b buffer.PacketBuffer) {
}

func (packet *PartyLeavePacket) Write(b buffer.PacketBuffer) {
}
