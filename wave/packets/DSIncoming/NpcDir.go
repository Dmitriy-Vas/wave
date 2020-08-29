package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type NpcDirPacket struct {
	*packets.DefaultPacket
	ID       int32 // NPC ID
	Dir      byte
	Position buffer.Vector2
}

func (n *NpcDirPacket) Read(b buffer.PacketBuffer) {
	n.ID = b.ReadInt(b.Bytes(), b.Index())
	n.Dir = b.ReadByte(b.Bytes(), b.Index())
	n.Position = b.ReadVector2(b.Bytes(), b.Index())
}

func (n *NpcDirPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), n.ID, b.Index())
	b.WriteByte(b.Bytes(), n.Dir, b.Index())
	b.WriteVector2(b.Bytes(), n.Position, b.Index())
}
