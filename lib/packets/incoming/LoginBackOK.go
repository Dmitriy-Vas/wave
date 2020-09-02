package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type LoginBackOKPacket struct {
	*wave.DefaultPacket
}

func (packet *LoginBackOKPacket) Read(b buffer.PacketBuffer) {
}

func (packet *LoginBackOKPacket) Write(b buffer.PacketBuffer) {
}
