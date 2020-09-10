package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *InteractionPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *InteractionPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *InteractionPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *InteractionPacket) SetSend(value bool) {
	d.Send = value
}

type InteractionPacket struct {
	ID      int64
	Send    bool
	IsClick bool
	X       int32
	Y       int32
}

func (packet *InteractionPacket) Read(b buffer.PacketBuffer) {
	packet.IsClick = b.ReadBool(b.Bytes(), b.Index())
	packet.X = b.ReadInt(b.Bytes(), b.Index())
	packet.Y = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *InteractionPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.IsClick, b.Index())
	b.WriteInt(b.Bytes(), packet.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Y, b.Index())
}
