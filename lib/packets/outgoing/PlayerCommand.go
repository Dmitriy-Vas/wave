package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerCommandPacket struct {
	*wave.DefaultPacket
	Command  string
	Emoji    int32
	Variable string
}

func (packet *PlayerCommandPacket) Read(b buffer.PacketBuffer) {
	packet.Command = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Emoji = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *PlayerCommandPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Command, b.Index())
	b.WriteInt(b.Bytes(), packet.Emoji, b.Index())
	b.WriteString(b.Bytes(), packet.Variable, b.Index())
}
