package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type BuyItemPacket struct {
	*wave.DefaultPacket
	IsCash    bool
	ShopSlot  int32
	Quantity  int32
	Variable5 byte
}

func (packet *BuyItemPacket) Read(b buffer.PacketBuffer) {
	packet.IsCash = b.ReadBool(b.Bytes(), b.Index())
	packet.ShopSlot = b.ReadInt(b.Bytes(), b.Index())
	packet.Quantity = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *BuyItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.IsCash, b.Index())
	b.WriteInt(b.Bytes(), packet.ShopSlot, b.Index())
	b.WriteInt(b.Bytes(), packet.Quantity, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable5, b.Index())
}
