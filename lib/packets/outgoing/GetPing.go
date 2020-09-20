package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *GetPingPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *GetPingPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *GetPingPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *GetPingPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type GetPingPacket struct {
	ID   int64
	Send bool
}

func (packet *GetPingPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *GetPingPacket) Write(_ buffer.PacketBuffer) {
}
