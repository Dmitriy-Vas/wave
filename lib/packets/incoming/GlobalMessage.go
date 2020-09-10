package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *GlobalMessagePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *GlobalMessagePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *GlobalMessagePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *GlobalMessagePacket) SetSend(value bool) {
	d.Send = value
}

type GlobalMessagePacket struct {
	ID    int64
	Send  bool
	Type  byte
	Num   int32
	Text  string
	Color string
}

func (packet *GlobalMessagePacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *GlobalMessagePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteString(b.Bytes(), packet.Color, b.Index())
}
