package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type UrlPacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 string
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &UrlPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *UrlPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *UrlPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteString(b.Bytes(), packet.Variable1, b.Index())
}
