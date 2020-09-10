package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *QuestMsgPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *QuestMsgPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *QuestMsgPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *QuestMsgPacket) SetSend(value bool) {
	d.Send = value
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
