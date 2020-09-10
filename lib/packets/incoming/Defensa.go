package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *DefensaPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *DefensaPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *DefensaPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *DefensaPacket) SetSend(value bool) {
	d.Send = value
}

type DefensaPacket struct {
	ID        int64
	Send      bool
	PlayerNum int32
	ShieldHP  int64
	Num2      int64
	Variable4 int64
}

func (packet *DefensaPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.ShieldHP = b.ReadLong(b.Bytes(), b.Index())
	packet.Num2 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *DefensaPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteLong(b.Bytes(), packet.ShieldHP, b.Index())
	b.WriteLong(b.Bytes(), packet.Num2, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable4, b.Index())
}
