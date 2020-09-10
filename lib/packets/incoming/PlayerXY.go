package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerXYPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerXYPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerXYPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerXYPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerXYPacket struct {
	ID   int64
	Send bool
	X    int32
	Y    int32
	Dir  byte
}

func (p *PlayerXYPacket) Read(b buffer.PacketBuffer) {
	p.X = b.ReadInt(b.Bytes(), b.Index())
	p.Y = b.ReadInt(b.Bytes(), b.Index())
	p.Dir = b.ReadByte(b.Bytes(), b.Index())
}

func (p *PlayerXYPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), p.X, b.Index())
	b.WriteInt(b.Bytes(), p.Y, b.Index())
	b.WriteByte(b.Bytes(), p.Dir, b.Index())
}
