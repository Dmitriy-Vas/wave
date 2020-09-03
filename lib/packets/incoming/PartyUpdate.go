package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type PartyUpdatePacket struct {
	*wave.DefaultPacket
	Num   int32
	Party lib.PartyRec
}

func (packet *PartyUpdatePacket) Read(b buffer.PacketBuffer) {
	if packet.Num = b.ReadInt(b.Bytes(), b.Index()); packet.Num > 0 {
		party := lib.PartyRec{
			Leader:      b.ReadInt(b.Bytes(), b.Index()),
			Type:        b.ReadByte(b.Bytes(), b.Index()),
			MemberCount: b.ReadByte(b.Bytes(), b.Index()),
		}
		party.Member = make([]lib.MemberPartyRec, party.MemberCount)
		for i := range party.Member {
			party.Member[i].Index = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].Name = b.ReadString(b.Bytes(), b.Index(), 0)
			party.Member[i].Map = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].Classes = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].Level = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].Sprite = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].Hair = b.ReadInt(b.Bytes(), b.Index())
			party.Member[i].HairTint = b.ReadString(b.Bytes(), b.Index(), 0)
		}
		party.ShareItems = b.ReadBool(b.Bytes(), b.Index())
		packet.Party = party
	}
}

func (packet *PartyUpdatePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	if packet.Num > 0 {
		b.WriteInt(b.Bytes(), packet.Party.Leader, b.Index())
		b.WriteByte(b.Bytes(), packet.Party.Type, b.Index())
		b.WriteByte(b.Bytes(), packet.Party.MemberCount, b.Index())
		for _, member := range packet.Party.Member {
			b.WriteInt(b.Bytes(), member.Index, b.Index())
			b.WriteString(b.Bytes(), member.Name, b.Index())
			b.WriteInt(b.Bytes(), member.Map, b.Index())
			b.WriteInt(b.Bytes(), member.Classes, b.Index())
			b.WriteInt(b.Bytes(), member.Level, b.Index())
			b.WriteInt(b.Bytes(), member.Sprite, b.Index())
			b.WriteInt(b.Bytes(), member.Hair, b.Index())
			b.WriteString(b.Bytes(), member.HairTint, b.Index())
		}
		b.WriteBool(b.Bytes(), packet.Party.ShareItems, b.Index())
	}
}
