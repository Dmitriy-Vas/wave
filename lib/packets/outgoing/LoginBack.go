package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type LoginBackPacket struct {
	*wave.DefaultPacket
}

func (packet *LoginBackPacket) Read(b buffer.PacketBuffer) {
}

func (packet *LoginBackPacket) Write(b buffer.PacketBuffer) {
}
