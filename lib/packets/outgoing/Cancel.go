package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type CancelPacket struct {
	*wave.DefaultPacket
}

func (packet *CancelPacket) Read(b buffer.PacketBuffer) {
}

func (packet *CancelPacket) Write(b buffer.PacketBuffer) {
}
