package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerAwardsPacket struct {
	*wave.DefaultPacket
	AwardNum int32
	AP       int32
	Level    int32
	Count    int32
	GetDate  string
	WithLog  bool
}

func (packet *PlayerAwardsPacket) Read(b buffer.PacketBuffer) {
	packet.AwardNum = b.ReadInt(b.Bytes(), b.Index())
	packet.AP = b.ReadInt(b.Bytes(), b.Index())
	packet.Level = b.ReadInt(b.Bytes(), b.Index())
	packet.Count = b.ReadInt(b.Bytes(), b.Index())
	packet.GetDate = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.WithLog = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *PlayerAwardsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.AwardNum, b.Index())
	b.WriteInt(b.Bytes(), packet.AP, b.Index())
	b.WriteInt(b.Bytes(), packet.Level, b.Index())
	b.WriteInt(b.Bytes(), packet.Count, b.Index())
	b.WriteString(b.Bytes(), packet.GetDate, b.Index())
	b.WriteBool(b.Bytes(), packet.WithLog, b.Index())
}
