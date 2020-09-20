package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SpellAnimPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SpellAnimPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SpellAnimPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SpellAnimPacket) SetSend(value bool) {
	packet.Send = value
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
