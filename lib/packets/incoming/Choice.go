package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
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
	ID          int64
	Send        bool
	IsCondition bool
	Choice      lib.ChoiceRec
	Variable0   bool
	Text        string
}

func (packet *ChoicePacket) Read(b buffer.PacketBuffer) {
	packet.IsCondition = b.ReadBool(b.Bytes(), b.Index())
	if packet.IsCondition {
		packet.Choice = lib.ChoiceRec{
			Max:           b.ReadInt(b.Bytes(), b.Index()),
			ChoiceText:    b.ReadInt(b.Bytes(), b.Index()),
			ChoiceSelText: b.ReadInt(b.Bytes(), b.Index()),
			Code:          b.ReadBool(b.Bytes(), b.Index()),
		}
		packet.Choice.ChoiceArray = make([]string, packet.Choice.Max)
		for i := range packet.Choice.ChoiceArray {
			packet.Choice.ChoiceArray[i] = b.ReadString(b.Bytes(), b.Index(), 0)
		}
	} else {
		packet.Choice = lib.ChoiceRec{
			DialogType: b.ReadByte(b.Bytes(), b.Index()),
		}
		packet.Variable0 = b.ReadBool(b.Bytes(), b.Index())
		packet.Choice.ChoiceText = b.ReadInt(b.Bytes(), b.Index())
		packet.Choice.ChoiceSelText = b.ReadInt(b.Bytes(), b.Index())

		packet.Choice.Sprite = b.ReadInt(b.Bytes(), b.Index())
		packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	}
}

func (packet *ChoicePacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.IsCondition, b.Index())
	if packet.IsCondition {
		b.WriteInt(b.Bytes(), packet.Choice.Max, b.Index())
		b.WriteInt(b.Bytes(), packet.Choice.ChoiceText, b.Index())
		b.WriteInt(b.Bytes(), packet.Choice.ChoiceSelText, b.Index())
		b.WriteBool(b.Bytes(), packet.Choice.Code, b.Index())
		for _, choice := range packet.Choice.ChoiceArray {
			b.WriteString(b.Bytes(), choice, b.Index())
		}
	} else {
		b.WriteByte(b.Bytes(), packet.Choice.DialogType, b.Index())
		b.WriteBool(b.Bytes(), packet.Variable0, b.Index())
		b.WriteInt(b.Bytes(), packet.Choice.ChoiceText, b.Index())
		b.WriteInt(b.Bytes(), packet.Choice.ChoiceSelText, b.Index())
		b.WriteInt(b.Bytes(), packet.Choice.Sprite, b.Index())
		b.WriteString(b.Bytes(), packet.Text, b.Index())
	}
}
