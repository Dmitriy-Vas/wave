package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ClearSpellBufferPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ClearSpellBufferPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ClearSpellBufferPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ClearSpellBufferPacket) SetSend(value bool) {
	d.Send = value
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
