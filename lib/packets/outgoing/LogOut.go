package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *LogOutPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *LogOutPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *LogOutPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *LogOutPacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type LogOutPacket struct {
	ID   int64
	Send bool
}

func (packet *LogOutPacket) Read(b buffer.PacketBuffer) {
}

func (packet *LogOutPacket) Write(b buffer.PacketBuffer) {
}
