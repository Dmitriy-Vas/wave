package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *CloseShopPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CloseShopPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CloseShopPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CloseShopPacket) SetSend(value bool) {
	d.Send = value
}

type CloseShopPacket struct {
	ID    int64
	Send  bool
	Shop  bool
	Forge bool
}

func (packet *CloseShopPacket) Read(b buffer.PacketBuffer) {
	packet.Shop = b.ReadBool(b.Bytes(), b.Index())
	packet.Forge = b.ReadBool(b.Bytes(), b.Index())

}

func (packet *CloseShopPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.Shop, b.Index())
	b.WriteBool(b.Bytes(), packet.Forge, b.Index())
}
