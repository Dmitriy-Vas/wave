package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SteamAchievementPacket struct {
	*wave.DefaultPacket
	Achievement string
	Amount      string
}

func (packet *SteamAchievementPacket) Read(b buffer.PacketBuffer) {
	packet.Achievement = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Amount = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *SteamAchievementPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Achievement, b.Index())
	b.WriteString(b.Bytes(), packet.Amount, b.Index())
}
