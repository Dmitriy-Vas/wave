package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PingPacket struct {
	*wave.DefaultPacket
}

func (packet *PingPacket) Read(b buffer.PacketBuffer) {
}

func (packet *PingPacket) Write(b buffer.PacketBuffer) {
}
