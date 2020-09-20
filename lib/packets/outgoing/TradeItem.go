package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *TradeItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *TradeItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *TradeItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *TradeItemPacket) SetSend(value bool) {
	packet.Send = value
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
