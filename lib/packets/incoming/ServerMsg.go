package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ServerMsgPacket struct {
	*wave.DefaultPacket
	Lang  int32
	Text  string
	Color string
}

func (packet *ServerMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Lang = b.ReadInt(b.Bytes(), b.Index())
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *ServerMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Lang, b.Index())
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteString(b.Bytes(), packet.Color, b.Index())
}
