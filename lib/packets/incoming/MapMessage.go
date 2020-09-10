package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *MapMessagePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *MapMessagePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *MapMessagePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *MapMessagePacket) SetSend(value bool) {
	d.Send = value
}

type MapMessagePacket struct {
	ID      int64
	Send    bool
	Message string
}

func (packet *MapMessagePacket) Read(b buffer.PacketBuffer) {
	packet.Message = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *MapMessagePacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Message, b.Index())
}
