package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type WhosOnlinePacket struct {
	*wave.DefaultPacket
}

func (packet *WhosOnlinePacket) Read(b buffer.PacketBuffer) {
}

func (packet *WhosOnlinePacket) Write(b buffer.PacketBuffer) {
}
