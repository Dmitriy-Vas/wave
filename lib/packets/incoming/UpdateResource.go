package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdateResourcePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdateResourcePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdateResourcePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdateResourcePacket) SetSend(value bool) {
	d.Send = value
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
