package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerCommandPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerCommandPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerCommandPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerCommandPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerCommandPacket struct {
	ID       int64
	Send     bool
	Command  string
	Emoji    int32
	Variable string
}

func (packet *PlayerCommandPacket) Read(b buffer.PacketBuffer) {
	packet.Command = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Emoji = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *PlayerCommandPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Command, b.Index())
	b.WriteInt(b.Bytes(), packet.Emoji, b.Index())
	b.WriteString(b.Bytes(), packet.Variable, b.Index())
}
