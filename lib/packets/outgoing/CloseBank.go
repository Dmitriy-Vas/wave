package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *CloseBankPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *CloseBankPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *CloseBankPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *CloseBankPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type CloseBankPacket struct {
	ID   int64
	Send bool
}

func (packet *CloseBankPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *CloseBankPacket) Write(_ buffer.PacketBuffer) {
}
