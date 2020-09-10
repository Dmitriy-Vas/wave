package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerCurrentPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerCurrentPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerCurrentPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerCurrentPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerCurrentPacket struct {
	ID            int64
	Send          bool
	PlayerNum     int32
	Pin           int32
	ProjectilSkin byte
	DamageSkin    byte
}

func (packet *PlayerCurrentPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Pin = b.ReadInt(b.Bytes(), b.Index())
	packet.ProjectilSkin = b.ReadByte(b.Bytes(), b.Index())
	packet.DamageSkin = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerCurrentPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Pin, b.Index())
	b.WriteByte(b.Bytes(), packet.ProjectilSkin, b.Index())
	b.WriteByte(b.Bytes(), packet.DamageSkin, b.Index())
}
