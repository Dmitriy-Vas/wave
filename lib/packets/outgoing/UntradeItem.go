package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UntradeItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UntradeItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UntradeItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UntradeItemPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type UntradeItemPacket struct {
	ID   int64
	Send bool
}

func (packet *UntradeItemPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *UntradeItemPacket) Write(_ buffer.PacketBuffer) {
}
