package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type QuestMsgPacket struct {
	*wave.DefaultPacket
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
