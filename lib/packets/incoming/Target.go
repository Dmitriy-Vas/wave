package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *TargetPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *TargetPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *TargetPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *TargetPacket) SetSend(value bool) {
	packet.Send = value
}

type TargetPacket struct {
	ID         int64
	Send       bool
	Target     int32
	TargetType byte
}

func (packet *TargetPacket) Read(b buffer.PacketBuffer) {
	packet.Target = b.ReadInt(b.Bytes(), b.Index())
	packet.TargetType = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *TargetPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Target, b.Index())
	b.WriteByte(b.Bytes(), packet.TargetType, b.Index())
}
