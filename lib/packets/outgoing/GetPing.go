package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type GetPingPacket struct {
	*wave.DefaultPacket
}

func (packet *GetPingPacket) Read(b buffer.PacketBuffer) {
}

func (packet *GetPingPacket) Write(b buffer.PacketBuffer) {
}
