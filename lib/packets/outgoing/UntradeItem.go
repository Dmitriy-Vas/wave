package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UntradeItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UntradeItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UntradeItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UntradeItemPacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type UntradeItemPacket struct {
	ID   int64
	Send bool
}

func (packet *UntradeItemPacket) Read(b buffer.PacketBuffer) {
}

func (packet *UntradeItemPacket) Write(b buffer.PacketBuffer) {
}
