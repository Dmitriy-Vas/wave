package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdatePlayerListPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdatePlayerListPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdatePlayerListPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdatePlayerListPacket) SetSend(value bool) {
	d.Send = value
}

type UpdatePlayerListPacket struct {
	ID        int64
	Send      bool
	Index     int32
	Variable1 int32
	Type      int32
}

func (packet *UpdatePlayerListPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Type = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *UpdatePlayerListPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Type, b.Index())
}
