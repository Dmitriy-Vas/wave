package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
// GetID returns packet ID.
func (d *RequestLevelUpPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *RequestLevelUpPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *RequestLevelUpPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *RequestLevelUpPacket) SetSend(value bool) {
	d.Send = value
}

type RequestLevelUpPacket struct {
	ID   int64
	Send bool
}

func (packet *RequestLevelUpPacket) Read(b buffer.PacketBuffer) {
}

func (packet *RequestLevelUpPacket) Write(b buffer.PacketBuffer) {
}
