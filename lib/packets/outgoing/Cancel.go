package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *CancelPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CancelPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CancelPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CancelPacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type CancelPacket struct {
	ID   int64
	Send bool
}

func (packet *CancelPacket) Read(b buffer.PacketBuffer) {
}

func (packet *CancelPacket) Write(b buffer.PacketBuffer) {
}
