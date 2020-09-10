package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *TradeItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *TradeItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *TradeItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *TradeItemPacket) SetSend(value bool) {
	d.Send = value
}

type TradeItemPacket struct {
	ID      int64
	Send    bool
	InvSlot byte
	Amount  int64
	IsSkull bool
}

func (packet *TradeItemPacket) Read(b buffer.PacketBuffer) {
	packet.InvSlot = b.ReadByte(b.Bytes(), b.Index())
	packet.Amount = b.ReadLong(b.Bytes(), b.Index())
	packet.IsSkull = b.ReadBool(b.Bytes(), b.Index())

}

func (packet *TradeItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvSlot, b.Index())
	b.WriteLong(b.Bytes(), packet.Amount, b.Index())
	b.WriteBool(b.Bytes(), packet.IsSkull, b.Index())
}
