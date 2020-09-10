package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ConstantDataPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ConstantDataPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ConstantDataPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ConstantDataPacket) SetSend(value bool) {
	d.Send = value
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
