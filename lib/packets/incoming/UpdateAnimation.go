package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateAnimationPacket struct {
	*wave.DefaultPacket
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
