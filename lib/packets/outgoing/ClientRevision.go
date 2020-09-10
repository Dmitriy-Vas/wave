package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ClientRevisionPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ClientRevisionPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ClientRevisionPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ClientRevisionPacket) SetSend(value bool) {
	d.Send = value
}

type ClientRevisionPacket struct {
	ID       int64
	Send     bool
	Unknown1 byte
	IsSteam  bool
}

func (c *ClientRevisionPacket) Read(b buffer.PacketBuffer) {
	c.Unknown1 = b.ReadByte(b.Bytes(), b.Index())
	c.IsSteam = b.ReadBool(b.Bytes(), b.Index())
}

func (c *ClientRevisionPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), c.Unknown1, b.Index())
	b.WriteBool(b.Bytes(), c.IsSteam, b.Index())
}
