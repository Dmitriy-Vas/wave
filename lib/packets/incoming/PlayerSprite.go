package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerSpritePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerSpritePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerSpritePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerSpritePacket) SetSend(value bool) {
	d.Send = value
}

type PlayerSpritePacket struct {
	ID        int64
	Send      bool
	PlayerNum int32
	Sprite    int32
	Polymorph int32
}

func (packet *PlayerSpritePacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Sprite = b.ReadInt(b.Bytes(), b.Index())
	packet.Polymorph = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerSpritePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Sprite, b.Index())
	b.WriteInt(b.Bytes(), packet.Polymorph, b.Index())
}
