package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerDirPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerDirPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerDirPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerDirPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerDirPacket struct {
	ID        int64
	Send      bool
	PlayerNum int32
	Dir       byte
}

func (packet *PlayerDirPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerDirPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
}
