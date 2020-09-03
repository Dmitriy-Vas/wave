package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type OpenBackgroundPacket struct {
	*wave.DefaultPacket
	Background int32
}

func (packet *OpenBackgroundPacket) Read(b buffer.PacketBuffer) {
	packet.Background = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *OpenBackgroundPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Background, b.Index())
}
