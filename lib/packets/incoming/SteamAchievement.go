package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SteamAchievementPacket struct {
	*wave.DefaultPacket
	Variable0   int64
	Achievement string
	Amount      string
}

func (packet *SteamAchievementPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Achievement = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Amount = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *SteamAchievementPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteString(b.Bytes(), packet.Achievement, b.Index())
	b.WriteString(b.Bytes(), packet.Amount, b.Index())
}
