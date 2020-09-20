package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *LogOutOKPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *LogOutOKPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *LogOutOKPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *LogOutOKPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type LogOutOKPacket struct {
	ID   int64
	Send bool
}

func (packet *LogOutOKPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *LogOutOKPacket) Write(_ buffer.PacketBuffer) {
}
