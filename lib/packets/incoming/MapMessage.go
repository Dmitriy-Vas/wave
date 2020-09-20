package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *MapMessagePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *MapMessagePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *MapMessagePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *MapMessagePacket) SetSend(value bool) {
	packet.Send = value
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
