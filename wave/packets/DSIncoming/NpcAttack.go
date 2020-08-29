package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type NpcAttackPacket struct {
	*packets.DefaultPacket
	ID        int32
	Damage    byte
	PlaySound bool

	Num   int32
	Vital int64
	HP    int64
}

func (n *NpcAttackPacket) Read(b buffer.PacketBuffer) {
	n.ID = b.ReadInt(b.Bytes(), b.Index())
	n.Damage = b.ReadByte(b.Bytes(), b.Index())
	n.PlaySound = b.ReadBool(b.Bytes(), b.Index())
	if n.ID > 0 {
		n.Num = b.ReadInt(b.Bytes(), b.Index())
		n.Vital = b.ReadLong(b.Bytes(), b.Index())
		n.HP = b.ReadLong(b.Bytes(), b.Index())
	}
}

func (n *NpcAttackPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), n.ID, b.Index())
	b.WriteByte(b.Bytes(), n.Damage, b.Index())
	b.WriteBool(b.Bytes(), n.PlaySound, b.Index())
	if n.ID > 0 {
		b.WriteInt(b.Bytes(), n.Num, b.Index())
		b.WriteLong(b.Bytes(), n.Vital, b.Index())
		b.WriteLong(b.Bytes(), n.HP, b.Index())
	}
}
