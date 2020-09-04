package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type AcceptTradePacket struct {
	*wave.DefaultPacket
}

func (packet *AcceptTradePacket) Read(b buffer.PacketBuffer) {
}

func (packet *AcceptTradePacket) Write(b buffer.PacketBuffer) {
}
