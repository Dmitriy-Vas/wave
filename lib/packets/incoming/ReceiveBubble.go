package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

// GetID returns packet ID.
func (packet *ReceiveBubblePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ReceiveBubblePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ReceiveBubblePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ReceiveBubblePacket) SetSend(value bool) {
	packet.Send = value
}

type ReceiveBubblePacket struct {
	ID        int64
	Send      bool
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
