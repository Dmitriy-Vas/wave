package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
// GetID returns packet ID.
func (d *InGamePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *InGamePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *InGamePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *InGamePacket) SetSend(value bool) {
	d.Send = value
}

type InGamePacket struct {
	ID   int64
	Send bool
}

func (packet *InGamePacket) Read(b buffer.PacketBuffer) {
}

func (packet *InGamePacket) Write(b buffer.PacketBuffer) {
}
