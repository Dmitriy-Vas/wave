package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type DeclineTradePacket struct {
	*wave.DefaultPacket
}

func (packet *DeclineTradePacket) Read(b buffer.PacketBuffer) {
}

func (packet *DeclineTradePacket) Write(b buffer.PacketBuffer) {
}
