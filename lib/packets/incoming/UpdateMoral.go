package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdateMoralPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdateMoralPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdateMoralPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdateMoralPacket) SetSend(value bool) {
	d.Send = value
}

type UpdateMoralPacket struct {
	ID        int64
	Send      bool
	Variable1 int32
}

func (packet *UpdateMoralPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	// TODO moral data
}

func (packet *UpdateMoralPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	// TODO moral data
}
