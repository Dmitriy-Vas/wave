package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *SendTimePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SendTimePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SendTimePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SendTimePacket) SetSend(value bool) {
	d.Send = value
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
