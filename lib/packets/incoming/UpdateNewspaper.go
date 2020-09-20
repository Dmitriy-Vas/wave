package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UpdateNewspaperPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UpdateNewspaperPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UpdateNewspaperPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UpdateNewspaperPacket) SetSend(value bool) {
	packet.Send = value
}

type UpdateNewspaperPacket struct {
	ID   int64
	Send bool
}

func (packet *UpdateNewspaperPacket) Read(_ buffer.PacketBuffer) {
	// TODO Newspaper data
}

func (packet *UpdateNewspaperPacket) Write(_ buffer.PacketBuffer) {
	// TODO Newspaper data
}
