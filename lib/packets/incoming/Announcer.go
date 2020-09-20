package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *AnnouncerPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *AnnouncerPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *AnnouncerPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *AnnouncerPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type AnnouncerPacket struct {
	ID   int64
	Send bool
}

func (packet *AnnouncerPacket) Read(_ buffer.PacketBuffer) {

}

func (packet *AnnouncerPacket) Write(_ buffer.PacketBuffer) {

}
