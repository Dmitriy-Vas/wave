package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (d *BuffInfoPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *BuffInfoPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *BuffInfoPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *BuffInfoPacket) SetSend(value bool) {
	d.Send = value
}

type BuffInfoPacket struct {
	ID         int64
	Send       bool
	Num        int32
	Variable2  bool
	Type       byte
	PlayerBuff lib.PlayerBuffsRec
}

func (packet *BuffInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadBool(b.Bytes(), b.Index())
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	if packet.Variable2 {
		packet.PlayerBuff = lib.PlayerBuffsRec{
			SpellNum: b.ReadInt(b.Bytes(), b.Index()),
			Timer:    b.ReadInt(b.Bytes(), b.Index()),
			Vital:    b.ReadInt(b.Bytes(), b.Index()),
		}
	}
}

func (packet *BuffInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable2, b.Index())
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	if packet.Variable2 {
		b.WriteInt(b.Bytes(), packet.PlayerBuff.SpellNum, b.Index())
		b.WriteInt(b.Bytes(), packet.PlayerBuff.Timer, b.Index())
		b.WriteInt(b.Bytes(), packet.PlayerBuff.Vital, b.Index())
	}
}
