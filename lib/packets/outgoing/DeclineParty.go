package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *DeclinePartyPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DeclinePartyPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DeclinePartyPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DeclinePartyPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type DeclinePartyPacket struct {
	ID   int64
	Send bool
}

func (packet *DeclinePartyPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *DeclinePartyPacket) Write(_ buffer.PacketBuffer) {
}
