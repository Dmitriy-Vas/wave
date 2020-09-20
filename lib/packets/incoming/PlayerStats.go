package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PlayerStatsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerStatsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerStatsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerStatsPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerStatsPacket struct {
	ID          int64
	Send        bool
	PlayerNum   int32
	PlayerStats []int64
	Points      int32
}

func (packet *PlayerStatsPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayerStats = make([]int64, 6) // TODO move to constants
	for i := range packet.PlayerStats {
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
