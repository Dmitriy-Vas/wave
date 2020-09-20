package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *LoginBackPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *LoginBackPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *LoginBackPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *LoginBackPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type LoginBackPacket struct {
	ID   int64
	Send bool
}

func (packet *LoginBackPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *LoginBackPacket) Write(_ buffer.PacketBuffer) {
}
