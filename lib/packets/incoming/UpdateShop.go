package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateShopPacket struct {
	*wave.DefaultPacket
	Num int32
}

func (packet *UpdateShopPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	// TODO Shop data
}

func (packet *UpdateShopPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	// TODO Shop data
}
