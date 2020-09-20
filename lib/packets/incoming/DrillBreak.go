package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *DrillBreakPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DrillBreakPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DrillBreakPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DrillBreakPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type DrillBreakPacket struct {
	ID   int64
	Send bool
}

func (packet *DrillBreakPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *DrillBreakPacket) Write(_ buffer.PacketBuffer) {
}
