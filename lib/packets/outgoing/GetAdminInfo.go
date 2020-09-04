package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type GetAdminInfoPacket struct {
	*wave.DefaultPacket
	Info int32
}

func (packet *GetAdminInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Info = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *GetAdminInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Info, b.Index())
}
