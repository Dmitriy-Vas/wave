package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *DefensaPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DefensaPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DefensaPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DefensaPacket) SetSend(value bool) {
	packet.Send = value
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
