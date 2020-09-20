package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *MapRespawnPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *MapRespawnPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *MapRespawnPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *MapRespawnPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type MapRespawnPacket struct {
	ID   int64
	Send bool
}

func (packet *MapRespawnPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *MapRespawnPacket) Write(_ buffer.PacketBuffer) {
}
