package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *DemoUpdateQuestPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DemoUpdateQuestPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DemoUpdateQuestPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DemoUpdateQuestPacket) SetSend(value bool) {
	packet.Send = value
}

type DemoUpdateQuestPacket struct {
	ID        int64
	Send      bool
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
