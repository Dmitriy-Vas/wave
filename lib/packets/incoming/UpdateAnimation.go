package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdateAnimationPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdateAnimationPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdateAnimationPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdateAnimationPacket) SetSend(value bool) {
	d.Send = value
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
