package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ProjectilPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ProjectilPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ProjectilPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ProjectilPacket) SetSend(value bool) {
	d.Send = value
}

type ProjectilPacket struct {
	ID           int64
	Send         bool
	ProjectilNum int32
	PlayerNum    int32
	Variable3    int32
	Combo        int32
	Dir          byte
}

func (packet *ProjectilPacket) Read(b buffer.PacketBuffer) {
	packet.ProjectilNum = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
	packet.Combo = b.ReadInt(b.Bytes(), b.Index())
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ProjectilPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ProjectilNum, b.Index())
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Combo, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
}
