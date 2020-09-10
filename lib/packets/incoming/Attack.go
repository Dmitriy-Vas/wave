package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *AttackPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *AttackPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *AttackPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *AttackPacket) SetSend(value bool) {
	d.Send = value
}

type AttackPacket struct {
	ID        int64
	Send      bool
	Player    int32
	Attacking byte
	FileID    byte
}

func (a *AttackPacket) Read(b buffer.PacketBuffer) {
	a.Player = b.ReadInt(b.Bytes(), b.Index())
	a.Attacking = b.ReadByte(b.Bytes(), b.Index())
	a.FileID = b.ReadByte(b.Bytes(), b.Index())
}

func (a *AttackPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), a.Player, b.Index())
	b.WriteByte(b.Bytes(), a.Attacking, b.Index())
	b.WriteByte(b.Bytes(), a.FileID, b.Index())
}
