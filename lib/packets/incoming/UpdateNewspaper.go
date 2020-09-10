package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdateNewspaperPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdateNewspaperPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdateNewspaperPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdateNewspaperPacket) SetSend(value bool) {
	d.Send = value
}

type UpdateNewspaperPacket struct {
	ID   int64
	Send bool
}

func (packet *UpdateNewspaperPacket) Read(b buffer.PacketBuffer) {
	// TODO Newspaper data
}

func (packet *UpdateNewspaperPacket) Write(b buffer.PacketBuffer) {
	// TODO Newspaper data
}
