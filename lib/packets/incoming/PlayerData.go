package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PlayerDataPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerDataPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerDataPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerDataPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerDataPacket struct {
	ID   int64
	Send bool
}

func (packet *PlayerDataPacket) Read(_ buffer.PacketBuffer) {

}

func (packet *PlayerDataPacket) Write(_ buffer.PacketBuffer) {

}
