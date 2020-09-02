package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerStatsPacket struct {
	*wave.DefaultPacket
	PlayerNum   int32
	PlayerStats []int64
	Points      int32
}

func (packet *PlayerStatsPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayerStats = make([]int64, 6) // TODO move to constants
	for i, _ := range packet.PlayerStats {
		packet.PlayerStats[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	packet.Points = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerStatsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	for _, stat := range packet.PlayerStats {
		b.WriteLong(b.Bytes(), stat, b.Index())
	}
	b.WriteInt(b.Bytes(), packet.Points, b.Index())
}
