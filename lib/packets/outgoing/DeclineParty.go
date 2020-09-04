package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type DeclinePartyPacket struct {
	*wave.DefaultPacket
}

func (packet *DeclinePartyPacket) Read(b buffer.PacketBuffer) {
}

func (packet *DeclinePartyPacket) Write(b buffer.PacketBuffer) {
}
