package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *WithdrawItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *WithdrawItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *WithdrawItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *WithdrawItemPacket) SetSend(value bool) {
	d.Send = value
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
