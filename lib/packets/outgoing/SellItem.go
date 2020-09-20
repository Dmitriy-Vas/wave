package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SellItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SellItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SellItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SellItemPacket) SetSend(value bool) {
	packet.Send = value
}

type SellItemPacket struct {
	ID       int64
	Send     bool
	InvSlot  byte
	Multiple int64
}

func (packet *SellItemPacket) Read(b buffer.PacketBuffer) {
	packet.InvSlot = b.ReadByte(b.Bytes(), b.Index())
	packet.Multiple = b.ReadLong(b.Bytes(), b.Index())

}

func (packet *SellItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvSlot, b.Index())
	b.WriteLong(b.Bytes(), packet.Multiple, b.Index())
}
