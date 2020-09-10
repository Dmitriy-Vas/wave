package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerMsgPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerMsgPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerMsgPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerMsgPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerMsgPacket struct {
	ID       int64
	Send     bool
	Text     string
	ToAll    int64
	MoveChat byte
}

func (packet *PlayerMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.ToAll = b.ReadLong(b.Bytes(), b.Index())
	packet.MoveChat = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *PlayerMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteLong(b.Bytes(), packet.ToAll, b.Index())
	b.WriteByte(b.Bytes(), packet.MoveChat, b.Index())
}
