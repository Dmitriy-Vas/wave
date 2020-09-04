package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

type PlayerSearchPacket struct {
	*wave.DefaultPacket
	TargetPos objects.Vector2
	Target    int32
	Type      byte
}

func (packet *PlayerSearchPacket) Read(b buffer.PacketBuffer) {
	packet.TargetPos = objects.Vector2{X: b.ReadInt(b.Bytes(), b.Index()), Y: b.ReadInt(b.Bytes(), b.Index())}
	packet.Target = b.ReadInt(b.Bytes(), b.Index())
	packet.Type = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *PlayerSearchPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.TargetPos.X, b.Index())
	b.WriteInt(b.Bytes(), packet.TargetPos.Y, b.Index())
	b.WriteInt(b.Bytes(), packet.Target, b.Index())
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
}
