package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PlayerMsgPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerMsgPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerMsgPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerMsgPacket) SetSend(value bool) {
	packet.Send = value
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
