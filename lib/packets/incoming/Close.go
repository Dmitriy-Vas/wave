package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ClosePacket struct {
	*wave.DefaultPacket
	Type byte
	Num  int32
	Text string
}

func (packet *ClosePacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *ClosePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteString(b.Bytes(), packet.Text, b.Index())
}
