package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type OpenBackgroundPacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 int32
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &OpenBackgroundPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *OpenBackgroundPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *OpenBackgroundPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
}
