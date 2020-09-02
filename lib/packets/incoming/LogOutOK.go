package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type LogOutOKPacket struct {
	*wave.DefaultPacket
}

func (packet *LogOutOKPacket) Read(b buffer.PacketBuffer) {
}

func (packet *LogOutOKPacket) Write(b buffer.PacketBuffer) {
}
