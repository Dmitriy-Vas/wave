package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type AnimationPacket struct {
	*wave.DefaultPacket
	Animation int32
	Duration  int32
	Position  objects.Vector2
	Prefix    string
	Type      byte
	Target    int32
	Sprite    int32
}

func (packet *AnimationPacket) Read(b buffer.PacketBuffer) {
	packet.Animation = b.ReadInt(b.Bytes(), b.Index())
	packet.Duration = b.ReadInt(b.Bytes(), b.Index())
	packet.Position = b.ReadVector2(b.Bytes(), b.Index())
	packet.Prefix = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Target = b.ReadInt(b.Bytes(), b.Index())
	packet.Sprite = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *AnimationPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Animation, b.Index())
	b.WriteInt(b.Bytes(), packet.Duration, b.Index())
	b.WriteVector2(b.Bytes(), packet.Position, b.Index())
	b.WriteString(b.Bytes(), packet.Prefix, b.Index())
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteInt(b.Bytes(), packet.Target, b.Index())
	b.WriteInt(b.Bytes(), packet.Sprite, b.Index())
}
