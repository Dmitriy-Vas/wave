package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ClosePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ClosePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ClosePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ClosePacket) SetSend(value bool) {
	d.Send = value
}

type ClosePacket struct {
	ID   int64
	Send bool
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
