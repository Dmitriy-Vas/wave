package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *BuyItemPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *BuyItemPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *BuyItemPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *BuyItemPacket) SetSend(value bool) {
	d.Send = value
}

type BuyItemPacket struct {
	ID        int64
	Send      bool
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
