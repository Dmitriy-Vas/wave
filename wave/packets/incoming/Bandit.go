package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type BanditPacket struct {
	*wave.DefaultPacket
	BanditSelect int64
}

func (packet *BanditPacket) Read(b buffer.PacketBuffer) {
	packet.BanditSelect = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *BanditPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.BanditSelect, b.Index())
}
