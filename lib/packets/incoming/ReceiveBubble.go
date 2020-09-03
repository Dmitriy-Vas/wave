package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

type ReceiveBubblePacket struct {
	*wave.DefaultPacket
	Target    int32
	Type      int32
	Variable3 bool
	Text      string
	Language  int32
	Color     string
	Pos       objects.Vector2
}

func (packet *ReceiveBubblePacket) Read(b buffer.PacketBuffer) {
	packet.Target = b.ReadInt(b.Bytes(), b.Index())
	packet.Type = b.ReadInt(b.Bytes(), b.Index())
	if packet.Variable3 = b.ReadBool(b.Bytes(), b.Index()); packet.Variable3 {
		packet.Language = b.ReadInt(b.Bytes(), b.Index())
	} else {
		packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)

	}
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Pos = objects.Vector2{
		X: b.ReadInt(b.Bytes(), b.Index()),
		Y: b.ReadInt(b.Bytes(), b.Index()),
	}
}

func (packet *ReceiveBubblePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Target, b.Index())
	b.WriteInt(b.Bytes(), packet.Type, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable3, b.Index())
	if packet.Variable3 {
		b.WriteInt(b.Bytes(), packet.Language, b.Index())
	} else {
		b.WriteString(b.Bytes(), packet.Text, b.Index())
	}
	b.WriteString(b.Bytes(), packet.Color, b.Index())
	b.WriteInt(b.Bytes(), packet.Pos.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Pos.Y, b.Index())
}
