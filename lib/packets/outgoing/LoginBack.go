package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *LoginBackPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *LoginBackPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *LoginBackPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *LoginBackPacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type LoginBackPacket struct {
	ID   int64
	Send bool
}

func (packet *LoginBackPacket) Read(b buffer.PacketBuffer) {
}

func (packet *LoginBackPacket) Write(b buffer.PacketBuffer) {
}
