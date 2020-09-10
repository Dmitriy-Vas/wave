package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
// GetID returns packet ID.
func (d *AcceptTradePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *AcceptTradePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *AcceptTradePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *AcceptTradePacket) SetSend(value bool) {
	d.Send = value
}

type AcceptTradePacket struct {
	ID   int64
	Send bool
}

func (packet *AcceptTradePacket) Read(b buffer.PacketBuffer) {
}

func (packet *AcceptTradePacket) Write(b buffer.PacketBuffer) {
}
