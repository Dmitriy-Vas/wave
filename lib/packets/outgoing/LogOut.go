package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *LogOutPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *LogOutPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *LogOutPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *LogOutPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type LogOutPacket struct {
	ID   int64
	Send bool
}

func (packet *LogOutPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *LogOutPacket) Write(_ buffer.PacketBuffer) {
}
