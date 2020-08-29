package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type PingPacket struct {
	*packets.DefaultPacket
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &PingPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *PingPacket) Read(b buffer.PacketBuffer) {
}

func (packet *PingPacket) Write(b buffer.PacketBuffer) {
}
