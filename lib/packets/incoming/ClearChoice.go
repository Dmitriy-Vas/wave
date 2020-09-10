package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ClearChoicePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ClearChoicePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ClearChoicePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ClearChoicePacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type ClearChoicePacket struct {
	ID   int64
	Send bool
}

func (packet *ClearChoicePacket) Read(b buffer.PacketBuffer) {
}

func (packet *ClearChoicePacket) Write(b buffer.PacketBuffer) {
}
