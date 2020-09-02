package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type NpcAttackPacket struct {
	*wave.DefaultPacket
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
