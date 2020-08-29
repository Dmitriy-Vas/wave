package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type BuddiesActionPacket struct {
	*packets.DefaultPacket
	Action   byte
	BuddieID byte
	Name     string
	Sex      byte
	Classes  byte

	FriendDate string
	Sprite     int32
	Hair       int32
	Paperdoll  string
	HairTint   string
	Online     bool
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &BuddiesActionPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *BuddiesActionPacket) Read(b buffer.PacketBuffer) {
	packet.Action = b.ReadByte(b.Bytes(), b.Index())
	packet.BuddieID = b.ReadByte(b.Bytes(), b.Index())
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Sex = b.ReadByte(b.Bytes(), b.Index())
	packet.Classes = b.ReadByte(b.Bytes(), b.Index())
	packet.FriendDate = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Sprite = b.ReadInt(b.Bytes(), b.Index())
	packet.Hair = b.ReadInt(b.Bytes(), b.Index())
	packet.Paperdoll = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.HairTint = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Online = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *BuddiesActionPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Action, b.Index())
	b.WriteByte(b.Bytes(), packet.BuddieID, b.Index())
	b.WriteString(b.Bytes(), packet.Name, b.Index())
	b.WriteByte(b.Bytes(), packet.Sex, b.Index())
	b.WriteByte(b.Bytes(), packet.Classes, b.Index())
	b.WriteString(b.Bytes(), packet.FriendDate, b.Index())
	b.WriteInt(b.Bytes(), packet.Sprite, b.Index())
	b.WriteInt(b.Bytes(), packet.Hair, b.Index())
	b.WriteString(b.Bytes(), packet.Paperdoll, b.Index())
	b.WriteString(b.Bytes(), packet.HairTint, b.Index())
	b.WriteBool(b.Bytes(), packet.Online, b.Index())
}
