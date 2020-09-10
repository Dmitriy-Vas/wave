package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ChoicePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ChoicePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ChoicePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ChoicePacket) SetSend(value bool) {
	d.Send = value
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
