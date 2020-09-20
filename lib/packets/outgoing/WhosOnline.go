package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *WhosOnlinePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *WhosOnlinePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *WhosOnlinePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *WhosOnlinePacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type WhosOnlinePacket struct {
	ID   int64
	Send bool
}

func (packet *WhosOnlinePacket) Read(_ buffer.PacketBuffer) {
}

func (packet *WhosOnlinePacket) Write(_ buffer.PacketBuffer) {
}
