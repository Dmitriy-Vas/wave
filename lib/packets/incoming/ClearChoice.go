package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ClearChoicePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ClearChoicePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ClearChoicePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ClearChoicePacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type ClearChoicePacket struct {
	ID   int64
	Send bool
}

func (packet *ClearChoicePacket) Read(_ buffer.PacketBuffer) {
}

func (packet *ClearChoicePacket) Write(_ buffer.PacketBuffer) {
}
