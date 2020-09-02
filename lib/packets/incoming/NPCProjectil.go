package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type NPCProjectilPacket struct {
	*wave.DefaultPacket
	NpcNum         int32
	Type           byte
	ProjectileNum  []int32
	NPCProjectiles []lib.NPCProjectilRec
	Direction      byte
}

func (packet *NPCProjectilPacket) Read(b buffer.PacketBuffer) {
	packet.NpcNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	switch packet.Type {
	case 0:
		packet.NPCProjectiles = make([]lib.NPCProjectilRec, 1)
		packet.ProjectileNum = make([]int32, 1)
	case 1:
		packet.NPCProjectiles = make([]lib.NPCProjectilRec, 4)
		packet.ProjectileNum = make([]int32, 4)
	}
	for i, _ := range packet.NPCProjectiles {
		packet.ProjectileNum[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.NPCProjectiles[i] = lib.NPCProjectilRec{
			Num:       int64(b.ReadInt(b.Bytes(), b.Index())),
			Direction: int64(b.ReadByte(b.Bytes(), b.Index())),
		}
	}
	packet.Direction = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *NPCProjectilPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.NpcNum, b.Index())
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	for i, _ := range packet.NPCProjectiles {
		b.WriteInt(b.Bytes(), packet.ProjectileNum[i], b.Index())
		b.WriteInt(b.Bytes(), int32(packet.NPCProjectiles[i].Num), b.Index())
		b.WriteByte(b.Bytes(), byte(packet.NPCProjectiles[i].Direction), b.Index())
	}
	b.WriteByte(b.Bytes(), packet.Direction, b.Index())
}
