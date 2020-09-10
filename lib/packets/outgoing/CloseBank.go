package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *CloseBankPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CloseBankPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CloseBankPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CloseBankPacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type CloseBankPacket struct {
	ID   int64
	Send bool
}

func (packet *CloseBankPacket) Read(b buffer.PacketBuffer) {
}

func (packet *CloseBankPacket) Write(b buffer.PacketBuffer) {
}
