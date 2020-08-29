package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

// Empty packet, only ID
type AnnouncerPacket struct {
	*packets.DefaultPacket
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &AnnouncerPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *AnnouncerPacket) Read(b buffer.PacketBuffer) {

}

func (packet *AnnouncerPacket) Write(b buffer.PacketBuffer) {

}
