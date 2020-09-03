package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type DrillBreakPacket struct {
	*wave.DefaultPacket
}

func (packet *DrillBreakPacket) Read(b buffer.PacketBuffer) {
}

func (packet *DrillBreakPacket) Write(b buffer.PacketBuffer) {
}
