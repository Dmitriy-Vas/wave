package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type MapRespawnPacket struct {
	*wave.DefaultPacket
}

func (packet *MapRespawnPacket) Read(b buffer.PacketBuffer) {
}

func (packet *MapRespawnPacket) Write(b buffer.PacketBuffer) {
}
