package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UpdateAnimationPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UpdateAnimationPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UpdateAnimationPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UpdateAnimationPacket) SetSend(value bool) {
	packet.Send = value
}

type UpdateAnimationPacket struct {
	ID           int64
	Send         bool
	AnimationNum int32
}

func (packet *UpdateAnimationPacket) Read(b buffer.PacketBuffer) {
	packet.AnimationNum = b.ReadInt(b.Bytes(), b.Index())
	// TODO animation data
}

func (packet *UpdateAnimationPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.AnimationNum, b.Index())
	// TODO animation data
}
