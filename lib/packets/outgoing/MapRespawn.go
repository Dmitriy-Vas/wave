package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *MapRespawnPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *MapRespawnPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *MapRespawnPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *MapRespawnPacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type MapRespawnPacket struct {
	ID   int64
	Send bool
}

func (packet *MapRespawnPacket) Read(b buffer.PacketBuffer) {
}

func (packet *MapRespawnPacket) Write(b buffer.PacketBuffer) {
}
