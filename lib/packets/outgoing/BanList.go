package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
// GetID returns packet ID.
func (d *BanListPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *BanListPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *BanListPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *BanListPacket) SetSend(value bool) {
	d.Send = value
}

type BanListPacket struct {
	ID   int64
	Send bool
}

func (packet *BanListPacket) Read(b buffer.PacketBuffer) {

}

func (packet *BanListPacket) Write(b buffer.PacketBuffer) {
}
