package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type InteractionPacket struct {
	*wave.DefaultPacket
	IsClick bool
	X       int32
	Y       int32
}

func (packet *InteractionPacket) Read(b buffer.PacketBuffer) {
	packet.IsClick = b.ReadBool(b.Bytes(), b.Index())
	packet.X = b.ReadInt(b.Bytes(), b.Index())
	packet.Y = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *InteractionPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.IsClick, b.Index())
	b.WriteInt(b.Bytes(), packet.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Y, b.Index())
}
