package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerInvitePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerInvitePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerInvitePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerInvitePacket) SetSend(value bool) {
	d.Send = value
}

type PlayerInvitePacket struct {
	ID   int64
	Send bool
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
