package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type AttackPacket struct {
	*packets.DefaultPacket
	ID        int32 // Player ID
	Attacking byte
	Arg       byte
}

func (a *AttackPacket) Read(b buffer.PacketBuffer) {
	a.ID = b.ReadInt(b.Bytes(), b.Index())
	a.Attacking = b.ReadByte(b.Bytes(), b.Index())
	a.Arg = b.ReadByte(b.Bytes(), b.Index())
}

func (a *AttackPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), a.ID, b.Index())
	b.WriteByte(b.Bytes(), a.Attacking, b.Index())
	b.WriteByte(b.Bytes(), a.Arg, b.Index())
}
