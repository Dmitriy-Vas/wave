package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ProjectilCrossPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ProjectilCrossPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ProjectilCrossPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ProjectilCrossPacket) SetSend(value bool) {
	d.Send = value
}

type ProjectilCrossPacket struct {
	ID             int64
	Send           bool
	PlayerNum      int64
	ProjectilNum   int64
	ProjectilIndex int64
	Combo          int64
	Direction      int64
}

func (packet *ProjectilCrossPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadLong(b.Bytes(), b.Index())
	packet.ProjectilNum = b.ReadLong(b.Bytes(), b.Index())
	packet.ProjectilIndex = b.ReadLong(b.Bytes(), b.Index())
	packet.Combo = b.ReadLong(b.Bytes(), b.Index())
	packet.Direction = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *ProjectilCrossPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteLong(b.Bytes(), packet.ProjectilNum, b.Index())
	b.WriteLong(b.Bytes(), packet.ProjectilIndex, b.Index())
	b.WriteLong(b.Bytes(), packet.Combo, b.Index())
	b.WriteLong(b.Bytes(), packet.Direction, b.Index())
}
