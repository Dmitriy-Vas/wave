package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *DeclineTradePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *DeclineTradePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *DeclineTradePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *DeclineTradePacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type DeclineTradePacket struct {
	ID   int64
	Send bool
}

func (packet *DeclineTradePacket) Read(b buffer.PacketBuffer) {
}

func (packet *DeclineTradePacket) Write(b buffer.PacketBuffer) {
}
