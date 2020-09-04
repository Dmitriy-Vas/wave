package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerMsgPacket struct {
	*wave.DefaultPacket
	Text     string
	ToAll    int64
	MoveChat byte
}

func (packet *PlayerMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.ToAll = b.ReadLong(b.Bytes(), b.Index())
	packet.MoveChat = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *PlayerMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteLong(b.Bytes(), packet.ToAll, b.Index())
	b.WriteByte(b.Bytes(), packet.MoveChat, b.Index())
}
