package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type CloseShopPacket struct {
	*wave.DefaultPacket
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
