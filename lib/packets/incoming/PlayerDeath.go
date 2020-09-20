package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PlayerDeathPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerDeathPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerDeathPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerDeathPacket) SetSend(value bool) {
	packet.Send = value
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
