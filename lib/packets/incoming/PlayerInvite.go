package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerInvitePacket struct {
	*wave.DefaultPacket
	Type byte
	Text string
}

func (packet *PlayerInvitePacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *PlayerInvitePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteString(b.Bytes(), packet.Text, b.Index())
}
