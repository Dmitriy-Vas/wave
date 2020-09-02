package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/buffer/objects"
)

type SpawnNPCPacket struct {
	*wave.DefaultPacket
	MapNpcNum     int32
	NpcNum        int32
	Pos           objects.Vector2
	Dir           byte
	HPSetTo       int64
	HighIndex     int32
	PlayAnimation bool
	Vital         []int32
}

func (packet *SpawnNPCPacket) Read(b buffer.PacketBuffer) {
	packet.MapNpcNum = b.ReadInt(b.Bytes(), b.Index())
	packet.NpcNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Pos = b.ReadVector2(b.Bytes(), b.Index())
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
	packet.HPSetTo = b.ReadLong(b.Bytes(), b.Index())
	packet.HighIndex = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayAnimation = b.ReadBool(b.Bytes(), b.Index())
	packet.Vital = make([]int32, 3) // TODO move to constants
	for i, _ := range packet.Vital {
		packet.Vital[i] = b.ReadInt(b.Bytes(), b.Index())
	}
}

func (packet *SpawnNPCPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MapNpcNum, b.Index())
	b.WriteInt(b.Bytes(), packet.NpcNum, b.Index())
	b.WriteVector2(b.Bytes(), packet.Pos, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
	b.WriteLong(b.Bytes(), packet.HPSetTo, b.Index())
	b.WriteInt(b.Bytes(), packet.HighIndex, b.Index())
	b.WriteBool(b.Bytes(), packet.PlayAnimation, b.Index())
	for _, vital := range packet.Vital {
		b.WriteInt(b.Bytes(), vital, b.Index())
	}
}
