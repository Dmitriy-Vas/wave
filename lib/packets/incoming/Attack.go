package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type AttackPacket struct {
	*wave.DefaultPacket
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
