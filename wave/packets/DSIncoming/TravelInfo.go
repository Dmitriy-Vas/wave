package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type TravelInfoPacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 bool
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &TravelInfoPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *TravelInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *TravelInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable1, b.Index())
}
