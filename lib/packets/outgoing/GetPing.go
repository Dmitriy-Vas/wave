package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
// GetID returns packet ID.
func (d *GetPingPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *GetPingPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *GetPingPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *GetPingPacket) SetSend(value bool) {
	d.Send = value
}

type GetPingPacket struct {
	ID   int64
	Send bool
}

func (packet *GetPingPacket) Read(b buffer.PacketBuffer) {
}

func (packet *GetPingPacket) Write(b buffer.PacketBuffer) {
}
