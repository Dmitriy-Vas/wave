package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type GlobalMessagePacket struct {
	*wave.DefaultPacket
	Type  byte
	ID    int32
	Text  string
	Color string
}

func (packet *GlobalMessagePacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.ID = b.ReadInt(b.Bytes(), b.Index())
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *GlobalMessagePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteInt(b.Bytes(), packet.ID, b.Index())
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteString(b.Bytes(), packet.Color, b.Index())
}
