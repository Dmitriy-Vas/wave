package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *DeclineTradePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DeclineTradePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DeclineTradePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DeclineTradePacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type DeclineTradePacket struct {
	ID   int64
	Send bool
}

func (packet *DeclineTradePacket) Read(_ buffer.PacketBuffer) {
}

func (packet *DeclineTradePacket) Write(_ buffer.PacketBuffer) {
}
