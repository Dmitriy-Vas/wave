package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *SpellAnimPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SpellAnimPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SpellAnimPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SpellAnimPacket) SetSend(value bool) {
	d.Send = value
}

type SpellAnimPacket struct {
	ID        int64
	Send      bool
	PlayerNum int32
	AnimNum   int32
}

func (packet *SpellAnimPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.AnimNum = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *SpellAnimPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.AnimNum, b.Index())
}
