package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdatePlayerHonorPacket struct {
	*wave.DefaultPacket
	Name   string
	Target int32
}

func (packet *UpdatePlayerHonorPacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Target = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *UpdatePlayerHonorPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
	b.WriteInt(b.Bytes(), packet.Target, b.Index())
}
