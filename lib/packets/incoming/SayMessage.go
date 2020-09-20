package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SayMessagePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SayMessagePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SayMessagePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SayMessagePacket) SetSend(value bool) {
	packet.Send = value
}

type SayMessagePacket struct {
	ID           int64
	Send         bool
	BubbleTarget int32
	Text         string
	Text2        string
	Color        string
	Escape       bool
	Variable6    byte
	Variable7    bool
}

func (packet *SayMessagePacket) Read(b buffer.PacketBuffer) {
	packet.BubbleTarget = b.ReadInt(b.Bytes(), b.Index())
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Text2 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Escape = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *SayMessagePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.BubbleTarget, b.Index())
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteString(b.Bytes(), packet.Text2, b.Index())
	b.WriteString(b.Bytes(), packet.Color, b.Index())
	b.WriteBool(b.Bytes(), packet.Escape, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable6, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable7, b.Index())
}
