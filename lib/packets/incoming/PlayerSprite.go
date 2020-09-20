package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PlayerSpritePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerSpritePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerSpritePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerSpritePacket) SetSend(value bool) {
	packet.Send = value
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
