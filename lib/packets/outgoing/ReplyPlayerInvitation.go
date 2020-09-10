package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ReplyPlayerInvitationPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ReplyPlayerInvitationPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ReplyPlayerInvitationPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ReplyPlayerInvitationPacket) SetSend(value bool) {
	d.Send = value
}

type ReplyPlayerInvitationPacket struct {
	ID     int64
	Send   bool
	Type   byte
	Action bool
}

func (packet *ReplyPlayerInvitationPacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Action = b.ReadBool(b.Bytes(), b.Index())

}

func (packet *ReplyPlayerInvitationPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteBool(b.Bytes(), packet.Action, b.Index())
}
