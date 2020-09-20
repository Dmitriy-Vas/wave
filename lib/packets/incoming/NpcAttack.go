package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *NpcAttackPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *NpcAttackPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *NpcAttackPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *NpcAttackPacket) SetSend(value bool) {
	packet.Send = value
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

func (packet *NpcAttackPacket) Read(b buffer.PacketBuffer) {
	packet.NpcNum = b.ReadInt(b.Bytes(), b.Index())
	packet.AttackNum = b.ReadByte(b.Bytes(), b.Index())
	packet.PlaySound = b.ReadBool(b.Bytes(), b.Index())
	if packet.NpcNum > 0 {
		packet.Num = b.ReadInt(b.Bytes(), b.Index())
		packet.Vital = b.ReadLong(b.Bytes(), b.Index())
		packet.HP = b.ReadLong(b.Bytes(), b.Index())
	}
}

func (packet *NpcAttackPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.NpcNum, b.Index())
	b.WriteByte(b.Bytes(), packet.AttackNum, b.Index())
	b.WriteBool(b.Bytes(), packet.PlaySound, b.Index())
	if packet.NpcNum > 0 {
		b.WriteInt(b.Bytes(), packet.Num, b.Index())
		b.WriteLong(b.Bytes(), packet.Vital, b.Index())
		b.WriteLong(b.Bytes(), packet.HP, b.Index())
	}
}
