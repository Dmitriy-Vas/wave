package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *CheckMapPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CheckMapPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CheckMapPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CheckMapPacket) SetSend(value bool) {
	d.Send = value
}

type CheckMapPacket struct {
	ID          int64
	Send        bool
	MapNum      int32
	MapRevision int32
	X           int32
	Y           int32
}

func (packet *CheckMapPacket) Read(b buffer.PacketBuffer) {
	packet.MapNum = b.ReadInt(b.Bytes(), b.Index())
	packet.MapRevision = b.ReadInt(b.Bytes(), b.Index())
	packet.X = b.ReadInt(b.Bytes(), b.Index())
	packet.Y = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *CheckMapPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MapNum, b.Index())
	b.WriteInt(b.Bytes(), packet.MapRevision, b.Index())
	b.WriteInt(b.Bytes(), packet.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Y, b.Index())
}
