package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SteamAchievementPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SteamAchievementPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SteamAchievementPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SteamAchievementPacket) SetSend(value bool) {
	packet.Send = value
}

type SteamAchievementPacket struct {
	ID              int64
	Send            bool
	Achievement     string
	AchievementStat string
}

func (packet *SteamAchievementPacket) Read(b buffer.PacketBuffer) {
	packet.Achievement = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.AchievementStat = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *SteamAchievementPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Achievement, b.Index())
	b.WriteString(b.Bytes(), packet.AchievementStat, b.Index())
}
