package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *BuyItemPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *BuyItemPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *BuyItemPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *BuyItemPacket) SetSend(value bool) {
	packet.Send = value
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
