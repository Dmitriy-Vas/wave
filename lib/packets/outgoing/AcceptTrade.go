package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *AcceptTradePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *AcceptTradePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *AcceptTradePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *AcceptTradePacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type AcceptTradePacket struct {
	ID   int64
	Send bool
}

func (packet *AcceptTradePacket) Read(_ buffer.PacketBuffer) {
}

func (packet *AcceptTradePacket) Write(_ buffer.PacketBuffer) {
}
