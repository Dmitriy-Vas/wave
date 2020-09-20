package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PlayerAwardsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerAwardsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerAwardsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerAwardsPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerAwardsPacket struct {
	ID       int64
	Send     bool
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
