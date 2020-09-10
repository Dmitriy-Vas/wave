package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
// GetID returns packet ID.
func (d *DrillBreakPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *DrillBreakPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *DrillBreakPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *DrillBreakPacket) SetSend(value bool) {
	d.Send = value
}

type DrillBreakPacket struct {
	ID   int64
	Send bool
}

func (packet *DrillBreakPacket) Read(b buffer.PacketBuffer) {
}

func (packet *DrillBreakPacket) Write(b buffer.PacketBuffer) {
}
