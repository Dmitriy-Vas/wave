package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ClearSpellBufferPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ClearSpellBufferPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ClearSpellBufferPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ClearSpellBufferPacket) SetSend(value bool) {
	packet.Send = value
}

type ClearSpellBufferPacket struct {
	ID       int64
	Send     bool
	SpellNum byte
}

func (packet *ClearSpellBufferPacket) Read(b buffer.PacketBuffer) {
	packet.SpellNum = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ClearSpellBufferPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.SpellNum, b.Index())
}
