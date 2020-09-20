package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *QuestMsgPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *QuestMsgPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *QuestMsgPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *QuestMsgPacket) SetSend(value bool) {
	packet.Send = value
}

type QuestMsgPacket struct {
	ID       int64
	Send     bool
	Text     string
	Color    string
	Language int32
}

func (packet *QuestMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Language = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *QuestMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteString(b.Bytes(), packet.Color, b.Index())
	b.WriteInt(b.Bytes(), packet.Language, b.Index())
}
