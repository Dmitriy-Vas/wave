package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SendTimePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SendTimePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SendTimePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SendTimePacket) SetSend(value bool) {
	packet.Send = value
}

type SendTimePacket struct {
	ID      int64
	Send    bool
	DayTime bool
}

func (packet *SendTimePacket) Read(b buffer.PacketBuffer) {
	packet.DayTime = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *SendTimePacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.DayTime, b.Index())
}
