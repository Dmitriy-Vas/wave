package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *BanListPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *BanListPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *BanListPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *BanListPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type BanListPacket struct {
	ID   int64
	Send bool
}

func (packet *BanListPacket) Read(_ buffer.PacketBuffer) {

}

func (packet *BanListPacket) Write(_ buffer.PacketBuffer) {
}
