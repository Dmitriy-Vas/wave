package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type CloseBankPacket struct {
	*wave.DefaultPacket
}

func (packet *CloseBankPacket) Read(b buffer.PacketBuffer) {
}

func (packet *CloseBankPacket) Write(b buffer.PacketBuffer) {
}
