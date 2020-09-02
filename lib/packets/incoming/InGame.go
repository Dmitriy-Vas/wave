package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type InGamePacket struct {
	*wave.DefaultPacket
}

func (packet *InGamePacket) Read(b buffer.PacketBuffer) {
}

func (packet *InGamePacket) Write(b buffer.PacketBuffer) {
}
