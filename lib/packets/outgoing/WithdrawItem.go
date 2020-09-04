package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type WithdrawItemPacket struct {
	*wave.DefaultPacket
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
