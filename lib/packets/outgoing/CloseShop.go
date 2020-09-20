package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *CloseShopPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *CloseShopPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *CloseShopPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *CloseShopPacket) SetSend(value bool) {
	packet.Send = value
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
