package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *WhosOnlinePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *WhosOnlinePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *WhosOnlinePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *WhosOnlinePacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type WhosOnlinePacket struct {
	ID   int64
	Send bool
}

func (packet *WhosOnlinePacket) Read(b buffer.PacketBuffer) {
}

func (packet *WhosOnlinePacket) Write(b buffer.PacketBuffer) {
}
