package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (d *CashShopDataPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CashShopDataPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CashShopDataPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CashShopDataPacket) SetSend(value bool) {
	d.Send = value
}

type CashShopDataPacket struct {
	ID      int64
	Send    bool
	ShopNum byte
	Item    lib.CashShopRec
}

func (packet *CashShopDataPacket) Read(b buffer.PacketBuffer) {
	packet.ShopNum = b.ReadByte(b.Bytes(), b.Index())
	packet.Item = lib.ReadCashShopData(b)
}

func (packet *CashShopDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.ShopNum, b.Index())
	lib.WriteCashShopData(b, packet.Item)
}
