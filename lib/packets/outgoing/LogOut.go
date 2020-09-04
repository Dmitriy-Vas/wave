package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type LogOutPacket struct {
	*wave.DefaultPacket
}

func (packet *LogOutPacket) Read(b buffer.PacketBuffer) {
}

func (packet *LogOutPacket) Write(b buffer.PacketBuffer) {
}
