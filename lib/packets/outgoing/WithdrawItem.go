package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *WithdrawItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *WithdrawItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *WithdrawItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *WithdrawItemPacket) SetSend(value bool) {
	packet.Send = value
}

type WithdrawItemPacket struct {
	ID       int64
	Send     bool
	BankSlot int32
	Amount   int64
}

func (packet *WithdrawItemPacket) Read(b buffer.PacketBuffer) {
	packet.BankSlot = b.ReadInt(b.Bytes(), b.Index())
	packet.Amount = b.ReadLong(b.Bytes(), b.Index())

}

func (packet *WithdrawItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.BankSlot, b.Index())
	b.WriteLong(b.Bytes(), packet.Amount, b.Index())
}
