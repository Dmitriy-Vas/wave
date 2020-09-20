package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PlayerXYPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerXYPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerXYPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerXYPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerXYPacket struct {
	ID   int64
	Send bool
	X    int32
	Y    int32
	Dir  byte
}

func (packet *PlayerXYPacket) Read(b buffer.PacketBuffer) {
	packet.X = b.ReadInt(b.Bytes(), b.Index())
	packet.Y = b.ReadInt(b.Bytes(), b.Index())
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerXYPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Y, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
}
