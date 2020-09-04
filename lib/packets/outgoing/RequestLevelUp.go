package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type RequestLevelUpPacket struct {
	*wave.DefaultPacket
}

func (packet *RequestLevelUpPacket) Read(b buffer.PacketBuffer) {
}

func (packet *RequestLevelUpPacket) Write(b buffer.PacketBuffer) {
}
