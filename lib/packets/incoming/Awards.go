package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (packet *AwardsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *AwardsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *AwardsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *AwardsPacket) SetSend(value bool) {
	packet.Send = value
}

type AwardsPacket struct {
	ID              int64
	Send            bool
	AwardID         int32
	Awards          []lib.AwardsRec
	PlayerFavAwards []int32
}

func (packet *AwardsPacket) Read(b buffer.PacketBuffer) {
	packet.AwardID = b.ReadInt(b.Bytes(), b.Index())
	if packet.AwardID != 0 {
		packet.Awards[0].Name = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.Awards[0].Desc = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.Awards[0].Blocked = b.ReadBool(b.Bytes(), b.Index())
		packet.Awards[0].Type = byte(b.ReadInt(b.Bytes(), b.Index()))
		packet.Awards[0].LevelCount = b.ReadInt(b.Bytes(), b.Index())
		packet.Awards[0].Level = make([]lib.AwardsLevelRec, packet.Awards[0].LevelCount)
		for i := 1; int32(i) < packet.Awards[0].LevelCount; i++ {
			packet.Awards[0].Level[i].Icon = b.ReadInt(b.Bytes(), b.Index())
			packet.Awards[0].Level[i].Required = b.ReadInt(b.Bytes(), b.Index())
			packet.Awards[0].Level[i].AP = b.ReadInt(b.Bytes(), b.Index())
			packet.Awards[0].Level[i].Calaveras = b.ReadInt(b.Bytes(), b.Index())
		}
	} else {
		packet.Awards = make([]lib.AwardsRec, 200) // TODO int to const
		for i := range packet.Awards {
			packet.Awards[i].Name = b.ReadString(b.Bytes(), b.Index(), 0)
			packet.Awards[i].ESName = b.ReadString(b.Bytes(), b.Index(), 0)
			packet.Awards[i].Desc = b.ReadString(b.Bytes(), b.Index(), 0)
			packet.Awards[i].Blocked = b.ReadBool(b.Bytes(), b.Index())
			packet.Awards[i].Type = byte(b.ReadInt(b.Bytes(), b.Index()))
			packet.Awards[i].LevelCount = b.ReadInt(b.Bytes(), b.Index())
			packet.Awards[i].Level = make([]lib.AwardsLevelRec, packet.Awards[i].LevelCount)
			for j := range packet.Awards[i].Level {
				packet.Awards[i].Level[j].Icon = b.ReadInt(b.Bytes(), b.Index())
				packet.Awards[i].Level[j].Required = b.ReadInt(b.Bytes(), b.Index())
				packet.Awards[i].Level[j].AP = b.ReadInt(b.Bytes(), b.Index())
				packet.Awards[i].Level[j].Calaveras = b.ReadInt(b.Bytes(), b.Index())
			}
		}
	}
	packet.PlayerFavAwards = make([]int32, 7) // TODO int to const
	for i := range packet.PlayerFavAwards {
		packet.PlayerFavAwards[i] = b.ReadInt(b.Bytes(), b.Index())
	}
}

func (packet *AwardsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.AwardID, b.Index())
	if packet.AwardID != 0 {
		b.WriteString(b.Bytes(), packet.Awards[0].Name, b.Index())
		b.WriteString(b.Bytes(), packet.Awards[0].Desc, b.Index())
		b.WriteBool(b.Bytes(), packet.Awards[0].Blocked, b.Index())
		b.WriteInt(b.Bytes(), int32(packet.Awards[0].Type), b.Index())
		b.WriteInt(b.Bytes(), packet.Awards[0].LevelCount, b.Index())
		for _, level := range packet.Awards[0].Level {
			b.WriteInt(b.Bytes(), level.Icon, b.Index())
			b.WriteInt(b.Bytes(), level.Required, b.Index())
			b.WriteInt(b.Bytes(), level.AP, b.Index())
			b.WriteInt(b.Bytes(), level.Calaveras, b.Index())
		}
	} else {
		for _, award := range packet.Awards {
			b.WriteString(b.Bytes(), award.Name, b.Index())
			b.WriteString(b.Bytes(), award.ESName, b.Index())
			b.WriteString(b.Bytes(), award.Desc, b.Index())
			b.WriteBool(b.Bytes(), award.Blocked, b.Index())
			b.WriteInt(b.Bytes(), int32(award.Type), b.Index())
			b.WriteInt(b.Bytes(), award.LevelCount, b.Index())
			for _, level := range award.Level {
				b.WriteInt(b.Bytes(), level.Icon, b.Index())
				b.WriteInt(b.Bytes(), level.Required, b.Index())
				b.WriteInt(b.Bytes(), level.AP, b.Index())
				b.WriteInt(b.Bytes(), level.Calaveras, b.Index())
			}
		}
	}
	for i := range packet.PlayerFavAwards {
		b.WriteInt(b.Bytes(), packet.PlayerFavAwards[i], b.Index())
	}
}
