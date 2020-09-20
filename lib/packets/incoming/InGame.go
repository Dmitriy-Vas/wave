package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *InGamePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *InGamePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *InGamePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *InGamePacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type InGamePacket struct {
	ID   int64
	Send bool
}

func (packet *InGamePacket) Read(_ buffer.PacketBuffer) {
}

func (packet *InGamePacket) Write(_ buffer.PacketBuffer) {
}
