package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *CharacterConfirmPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CharacterConfirmPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CharacterConfirmPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CharacterConfirmPacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type CharacterConfirmPacket struct {
	ID   int64
	Send bool
}

func (packet *CharacterConfirmPacket) Read(b buffer.PacketBuffer) {
}

func (packet *CharacterConfirmPacket) Write(b buffer.PacketBuffer) {
}
