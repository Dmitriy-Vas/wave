package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ReplyPlayerInvitationPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ReplyPlayerInvitationPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ReplyPlayerInvitationPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ReplyPlayerInvitationPacket) SetSend(value bool) {
	packet.Send = value
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
