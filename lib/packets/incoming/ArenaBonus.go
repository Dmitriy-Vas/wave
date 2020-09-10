package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ArenaBonusPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ArenaBonusPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ArenaBonusPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ArenaBonusPacket) SetSend(value bool) {
	d.Send = value
}

type ArenaBonusPacket struct {
	ID        int64
	Send      bool
	Variable1 bool
	Variable2 byte
}

func (packet *ArenaBonusPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ArenaBonusPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
}
