package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *NpcAttackPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *NpcAttackPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *NpcAttackPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *NpcAttackPacket) SetSend(value bool) {
	d.Send = value
}

type NpcAttackPacket struct {
	ID        int64
	Send      bool
	NpcNum    int32
	AttackNum byte
	PlaySound bool
	Num       int32
	Vital     int64
	HP        int64
}

func (n *NpcAttackPacket) Read(b buffer.PacketBuffer) {
	n.NpcNum = b.ReadInt(b.Bytes(), b.Index())
	n.AttackNum = b.ReadByte(b.Bytes(), b.Index())
	n.PlaySound = b.ReadBool(b.Bytes(), b.Index())
	if n.NpcNum > 0 {
		n.Num = b.ReadInt(b.Bytes(), b.Index())
		n.Vital = b.ReadLong(b.Bytes(), b.Index())
		n.HP = b.ReadLong(b.Bytes(), b.Index())
	}
}

func (n *NpcAttackPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), n.NpcNum, b.Index())
	b.WriteByte(b.Bytes(), n.AttackNum, b.Index())
	b.WriteBool(b.Bytes(), n.PlaySound, b.Index())
	if n.NpcNum > 0 {
		b.WriteInt(b.Bytes(), n.Num, b.Index())
		b.WriteLong(b.Bytes(), n.Vital, b.Index())
		b.WriteLong(b.Bytes(), n.HP, b.Index())
	}
}
