package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type AnimationPacket struct {
	*packets.DefaultPacket
	Animation int32
	Duration  int32
	Position  buffer.Vector2
	Prefix    string
	Type      byte
	Target    int32
	Sprite    int32
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &AnimationPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
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
