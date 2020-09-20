package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UpdateResourcePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UpdateResourcePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UpdateResourcePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UpdateResourcePacket) SetSend(value bool) {
	packet.Send = value
}

type UpdateResourcePacket struct {
	ID        int64
	Send      bool
	Variable0 int32
}

func (packet *UpdateResourcePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadInt(b.Bytes(), b.Index())
	// TODO resource data
}

func (packet *UpdateResourcePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable0, b.Index())
	// TODO resource data
}
