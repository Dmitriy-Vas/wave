package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

// GetID returns packet ID.
func (packet *SpawnNPCPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SpawnNPCPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SpawnNPCPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SpawnNPCPacket) SetSend(value bool) {
	packet.Send = value
}

type SpawnNPCPacket struct {
	ID            int64
	Send          bool
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
	packet.Pos = objects.Vector2{
		X: b.ReadInt(b.Bytes(), b.Index()),
		Y: b.ReadInt(b.Bytes(), b.Index()),
	}
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
	packet.HPSetTo = b.ReadLong(b.Bytes(), b.Index())
	packet.HighIndex = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayAnimation = b.ReadBool(b.Bytes(), b.Index())
	packet.Vital = make([]int32, 3) // TODO move to constants
	for i := range packet.Vital {
		packet.Vital[i] = b.ReadInt(b.Bytes(), b.Index())
	}
}

func (packet *SpawnNPCPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MapNpcNum, b.Index())
	b.WriteInt(b.Bytes(), packet.NpcNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Pos.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Pos.Y, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
	b.WriteLong(b.Bytes(), packet.HPSetTo, b.Index())
	b.WriteInt(b.Bytes(), packet.HighIndex, b.Index())
	b.WriteBool(b.Bytes(), packet.PlayAnimation, b.Index())
	for _, vital := range packet.Vital {
		b.WriteInt(b.Bytes(), vital, b.Index())
	}
}
