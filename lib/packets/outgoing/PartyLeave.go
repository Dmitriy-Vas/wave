package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PartyLeavePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PartyLeavePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PartyLeavePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PartyLeavePacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type PartyLeavePacket struct {
	ID   int64
	Send bool
}

func (packet *PartyLeavePacket) Read(_ buffer.PacketBuffer) {
}

func (packet *PartyLeavePacket) Write(_ buffer.PacketBuffer) {
}
