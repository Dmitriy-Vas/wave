package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerStatsPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerStatsPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerStatsPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerStatsPacket) SetSend(value bool) {
	d.Send = value
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
