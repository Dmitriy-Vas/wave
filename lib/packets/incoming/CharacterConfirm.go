package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *CharacterConfirmPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *CharacterConfirmPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *CharacterConfirmPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *CharacterConfirmPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type CharacterConfirmPacket struct {
	ID   int64
	Send bool
}

func (packet *CharacterConfirmPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *CharacterConfirmPacket) Write(_ buffer.PacketBuffer) {
}
