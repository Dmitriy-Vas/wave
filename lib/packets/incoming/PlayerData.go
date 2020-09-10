package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerDataPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerDataPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerDataPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerDataPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerDataPacket struct {
	ID   int64
	Send bool
}

func (packet *PlayerDataPacket) Read(b buffer.PacketBuffer) {

}

func (packet *PlayerDataPacket) Write(b buffer.PacketBuffer) {

}
