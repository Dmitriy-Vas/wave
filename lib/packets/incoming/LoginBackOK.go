package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *LoginBackOKPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *LoginBackOKPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *LoginBackOKPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *LoginBackOKPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type LoginBackOKPacket struct {
	ID   int64
	Send bool
}

func (packet *LoginBackOKPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *LoginBackOKPacket) Write(_ buffer.PacketBuffer) {
}
