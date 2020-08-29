package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type UpdateAnimationPacket struct {
	*packets.DefaultPacket
	Variable0 int64
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &UpdateAnimationPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *UpdateAnimationPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *UpdateAnimationPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
}