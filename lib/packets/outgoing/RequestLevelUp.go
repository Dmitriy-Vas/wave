package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *RequestLevelUpPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *RequestLevelUpPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *RequestLevelUpPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *RequestLevelUpPacket) SetSend(value bool) {
	packet.Send = value
}

// Empty packet
type RequestLevelUpPacket struct {
	ID   int64
	Send bool
}

func (packet *RequestLevelUpPacket) Read(_ buffer.PacketBuffer) {
}

func (packet *RequestLevelUpPacket) Write(_ buffer.PacketBuffer) {
}
