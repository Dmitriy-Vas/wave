package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ArenaBonusPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ArenaBonusPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ArenaBonusPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ArenaBonusPacket) SetSend(value bool) {
	packet.Send = value
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
