package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerXYMapPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerXYMapPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerXYMapPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerXYMapPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerXYMapPacket struct {
	ID    int64
	Send  bool
	Index int32
	X     int32
	Y     int32
	Dir   byte
}

func (packet *PlayerXYMapPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.X = b.ReadInt(b.Bytes(), b.Index())
	packet.Y = b.ReadInt(b.Bytes(), b.Index())
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerXYMapPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteInt(b.Bytes(), packet.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Y, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
}
