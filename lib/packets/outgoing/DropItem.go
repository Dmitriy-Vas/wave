package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type DropItemPacket struct {
	*wave.DefaultPacket
	InvNum byte
	Amount int64
}

func (packet *DropItemPacket) Read(b buffer.PacketBuffer) {
	packet.InvNum = b.ReadByte(b.Bytes(), b.Index())
	packet.Amount = b.ReadLong(b.Bytes(), b.Index())

}

func (packet *DropItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.InvNum, b.Index())
	b.WriteLong(b.Bytes(), packet.Amount, b.Index())
}
