package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ChoicePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ChoicePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ChoicePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ChoicePacket) SetSend(value bool) {
	packet.Send = value
}

type ChoicePacket struct {
	ID             int64
	Send           bool
	IsCondition    bool
	MyChoice       int32
	TextCode       string
	DialogType     byte
	ConditionLineX int32
	ConditionLineY int32
}

func (packet *ChoicePacket) Read(b buffer.PacketBuffer) {
	packet.IsCondition = b.ReadBool(b.Bytes(), b.Index())
	packet.MyChoice = b.ReadInt(b.Bytes(), b.Index())
	if packet.IsCondition {
		packet.TextCode = b.ReadString(b.Bytes(), b.Index(), 0)
	}
	packet.DialogType = b.ReadByte(b.Bytes(), b.Index())
	packet.ConditionLineX = b.ReadInt(b.Bytes(), b.Index())
	packet.ConditionLineY = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *ChoicePacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.IsCondition, b.Index())
	b.WriteInt(b.Bytes(), packet.MyChoice, b.Index())
	if packet.IsCondition {
		b.WriteString(b.Bytes(), packet.TextCode, b.Index())
	}
	b.WriteByte(b.Bytes(), packet.DialogType, b.Index())
	b.WriteInt(b.Bytes(), packet.ConditionLineX, b.Index())
	b.WriteInt(b.Bytes(), packet.ConditionLineY, b.Index())
}
