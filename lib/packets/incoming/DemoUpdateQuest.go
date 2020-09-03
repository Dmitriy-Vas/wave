package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type DemoUpdateQuestPacket struct {
	*wave.DefaultPacket
	QuestTasl int32
	ItemCount int32
	NpcCount  int32
	Num       int32
	Stage     bool
	Variable6 bool
}

func (packet *DemoUpdateQuestPacket) Read(b buffer.PacketBuffer) {
	packet.QuestTasl = b.ReadInt(b.Bytes(), b.Index())
	packet.ItemCount = b.ReadInt(b.Bytes(), b.Index())
	packet.NpcCount = b.ReadInt(b.Bytes(), b.Index())
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.Stage = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *DemoUpdateQuestPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.QuestTasl, b.Index())
	b.WriteInt(b.Bytes(), packet.ItemCount, b.Index())
	b.WriteInt(b.Bytes(), packet.NpcCount, b.Index())
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteBool(b.Bytes(), packet.Stage, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable6, b.Index())
}
