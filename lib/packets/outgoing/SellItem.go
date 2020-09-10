package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *SellItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SellItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SellItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SellItemPacket) SetSend(value bool) {
	d.Send = value
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
