package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *CancelPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *CancelPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *CancelPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *CancelPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type CancelPacket struct {
	ID   int64
	Send bool
}

func (packet *CancelPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *CancelPacket) Write(_ buffer.PacketBuffer) {
}
