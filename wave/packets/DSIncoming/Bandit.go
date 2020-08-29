package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type BanditPacket struct {
	*packets.DefaultPacket
	BanditSelect int64
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &BanditPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *BanditPacket) Read(b buffer.PacketBuffer) {
	packet.BanditSelect = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *BanditPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.BanditSelect, b.Index())
}
