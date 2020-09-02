package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/buffer/objects"
)

type NpcDeadPacket struct {
	*wave.DefaultPacket
	Index      int32
	Variable2  int32
	Animation  int32
	Pos        objects.Vector2
	EventCount int32
	Right      string // Name of player to compare
}

func (packet *NpcDeadPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Animation = b.ReadInt(b.Bytes(), b.Index())
	packet.Pos = b.ReadVector2(b.Bytes(), b.Index())
	packet.EventCount = b.ReadInt(b.Bytes(), b.Index())
	packet.Right = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *NpcDeadPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Animation, b.Index())
	b.WriteVector2(b.Bytes(), packet.Pos, b.Index())
	b.WriteInt(b.Bytes(), packet.EventCount, b.Index())
	b.WriteString(b.Bytes(), packet.Right, b.Index())
}
