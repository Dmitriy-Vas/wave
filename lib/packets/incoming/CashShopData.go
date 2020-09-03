package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type CashShopDataPacket struct {
	*wave.DefaultPacket
	Variable1 byte
}

func (packet *CashShopDataPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	// TODO cash data
}

func (packet *CashShopDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	// TODO cash data
}
