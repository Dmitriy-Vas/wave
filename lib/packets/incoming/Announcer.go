package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *AnnouncerPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *AnnouncerPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *AnnouncerPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *AnnouncerPacket) SetSend(value bool) {
	d.Send = value
}

// Empty packet
type AnnouncerPacket struct {
	ID   int64
	Send bool
}

func (packet *AnnouncerPacket) Read(b buffer.PacketBuffer) {

}

func (packet *AnnouncerPacket) Write(b buffer.PacketBuffer) {

}
