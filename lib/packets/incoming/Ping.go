package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PingPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PingPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PingPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PingPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type PingPacket struct {
	ID   int64
	Send bool
}

func (packet *PingPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *PingPacket) Write(_ buffer.PacketBuffer) {
}
