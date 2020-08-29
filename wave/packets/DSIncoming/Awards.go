package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type AwardsPacket struct {
	*packets.DefaultPacket
	AwardID         int32
	Awards          []Award
	PlayerFavAwards map[int]int32
}

type Award struct {
	Name        string
	ESName      string
	Description string
	Blocked     bool
	Type        int32
	LevelCount  int32
	AwardLevels []AwardLevel
}

type AwardLevel struct {
	Icon      int32
	Required  int32
	AP        int32
	Calaveras int32
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &AwardsPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *AwardsPacket) Read(b buffer.PacketBuffer) {
	packet.AwardID = b.ReadInt(b.Bytes(), b.Index())
	if packet.AwardID != 0 {
		packet.Awards[0].Name = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.Awards[0].Description = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.Awards[0].Blocked = b.ReadBool(b.Bytes(), b.Index())
		packet.Awards[0].Type = b.ReadInt(b.Bytes(), b.Index())
		packet.Awards[0].LevelCount = b.ReadInt(b.Bytes(), b.Index())
		packet.Awards[0].AwardLevels = make([]AwardLevel, packet.Awards[0].LevelCount)
		for i := 1; int32(i) < packet.Awards[0].LevelCount; i++ {
			packet.Awards[0].AwardLevels[i].Icon = b.ReadInt(b.Bytes(), b.Index())
			packet.Awards[0].AwardLevels[i].Required = b.ReadInt(b.Bytes(), b.Index())
			packet.Awards[0].AwardLevels[i].AP = b.ReadInt(b.Bytes(), b.Index())
			packet.Awards[0].AwardLevels[i].Calaveras = b.ReadInt(b.Bytes(), b.Index())
		}
	} else {
		packet.Awards = make([]Award, 200)
		for i := 1; i <= 200; i++ {
			packet.Awards[i].Name = b.ReadString(b.Bytes(), b.Index(), 0)
			packet.Awards[i].ESName = b.ReadString(b.Bytes(), b.Index(), 0)
			packet.Awards[i].Description = b.ReadString(b.Bytes(), b.Index(), 0)
			packet.Awards[i].Blocked = b.ReadBool(b.Bytes(), b.Index())
			packet.Awards[i].Type = b.ReadInt(b.Bytes(), b.Index())
			packet.Awards[i].LevelCount = b.ReadInt(b.Bytes(), b.Index())
			packet.Awards[i].AwardLevels = make([]AwardLevel, packet.Awards[i].LevelCount)
			for i := 1; int32(i) < packet.Awards[0].LevelCount; i++ {
				packet.Awards[i].AwardLevels[i].Icon = b.ReadInt(b.Bytes(), b.Index())
				packet.Awards[i].AwardLevels[i].Required = b.ReadInt(b.Bytes(), b.Index())
				packet.Awards[i].AwardLevels[i].AP = b.ReadInt(b.Bytes(), b.Index())
				packet.Awards[i].AwardLevels[i].Calaveras = b.ReadInt(b.Bytes(), b.Index())
			}
		}
	}
	packet.PlayerFavAwards = make(map[int]int32, 6)
	for i := 0; i <= 6; i++ {
		value := b.ReadInt(b.Bytes(), b.Index())
		packet.PlayerFavAwards[i] = value
	}
}

func (packet *AwardsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.AwardID, b.Index())
	if packet.AwardID != 0 {
		b.WriteString(b.Bytes(), packet.Awards[0].Name, b.Index())
		b.WriteString(b.Bytes(), packet.Awards[0].Description, b.Index())
		b.WriteBool(b.Bytes(), packet.Awards[0].Blocked, b.Index())
		b.WriteInt(b.Bytes(), packet.Awards[0].Type, b.Index())
		b.WriteInt(b.Bytes(), packet.Awards[0].LevelCount, b.Index())
		packet.Awards[0].AwardLevels = make([]AwardLevel, packet.Awards[0].LevelCount)
		for _, level := range packet.Awards[0].AwardLevels {
			b.WriteInt(b.Bytes(), level.Icon, b.Index())
			b.WriteInt(b.Bytes(), level.Required, b.Index())
			b.WriteInt(b.Bytes(), level.AP, b.Index())
			b.WriteInt(b.Bytes(), level.Calaveras, b.Index())
		}
	} else {
		packet.Awards = make([]Award, 200)
		for _, award := range packet.Awards {
			b.WriteString(b.Bytes(), award.Name, b.Index())
			b.WriteString(b.Bytes(), award.ESName, b.Index())
			b.WriteString(b.Bytes(), award.Description, b.Index())
			b.WriteBool(b.Bytes(), award.Blocked, b.Index())
			b.WriteInt(b.Bytes(), award.Type, b.Index())
			b.WriteInt(b.Bytes(), award.LevelCount, b.Index())
			for _, level := range award.AwardLevels {
				b.WriteInt(b.Bytes(), level.Icon, b.Index())
				b.WriteInt(b.Bytes(), level.Required, b.Index())
				b.WriteInt(b.Bytes(), level.AP, b.Index())
				b.WriteInt(b.Bytes(), level.Calaveras, b.Index())
			}
		}
	}
	for _, value := range packet.PlayerFavAwards {
		b.WriteInt(b.Bytes(), value, b.Index())
	}
}
