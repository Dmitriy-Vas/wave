package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
// GetID returns packet ID.
func (d *DeclinePartyPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *DeclinePartyPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *DeclinePartyPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *DeclinePartyPacket) SetSend(value bool) {
	d.Send = value
}

type DeclinePartyPacket struct {
	ID   int64
	Send bool
}

func (packet *DeclinePartyPacket) Read(b buffer.PacketBuffer) {
}

func (packet *DeclinePartyPacket) Write(b buffer.PacketBuffer) {
}
