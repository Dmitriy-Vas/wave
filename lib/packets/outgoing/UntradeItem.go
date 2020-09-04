package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type UntradeItemPacket struct {
	*wave.DefaultPacket
}

func (packet *UntradeItemPacket) Read(b buffer.PacketBuffer) {
}

func (packet *UntradeItemPacket) Write(b buffer.PacketBuffer) {
}
