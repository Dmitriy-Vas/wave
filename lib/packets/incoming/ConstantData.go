package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ConstantDataPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ConstantDataPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ConstantDataPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ConstantDataPacket) SetSend(value bool) {
	packet.Send = value
}

type ConstantDataPacket struct {
	ID         int64
	Send       bool
	MaxPlayers int32
	MaxLevels  int32
}

func (packet *ConstantDataPacket) Read(b buffer.PacketBuffer) {
	packet.MaxPlayers = b.ReadInt(b.Bytes(), b.Index())
	packet.MaxLevels = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *ConstantDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MaxPlayers, b.Index())
	b.WriteInt(b.Bytes(), packet.MaxLevels, b.Index())
}
