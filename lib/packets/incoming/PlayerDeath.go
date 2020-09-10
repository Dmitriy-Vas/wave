package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerDeathPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerDeathPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerDeathPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerDeathPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerDeathPacket struct {
	ID        int64
	Send      bool
	PlayerNum int32
	Ghost     byte
}

func (packet *PlayerDeathPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Ghost = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerDeathPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteByte(b.Bytes(), packet.Ghost, b.Index())
}
