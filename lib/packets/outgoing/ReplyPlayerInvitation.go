package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ReplyPlayerInvitationPacket struct {
	*wave.DefaultPacket
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
