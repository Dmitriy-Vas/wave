package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
// GetID returns packet ID.
func (d *PingPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PingPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PingPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PingPacket) SetSend(value bool) {
	d.Send = value
}

type PingPacket struct {
	ID   int64
	Send bool
}

func (packet *PingPacket) Read(b buffer.PacketBuffer) {
}

func (packet *PingPacket) Write(b buffer.PacketBuffer) {
}
