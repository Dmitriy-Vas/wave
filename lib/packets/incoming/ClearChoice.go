package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type ClearChoicePacket struct {
	*wave.DefaultPacket
}

func (packet *ClearChoicePacket) Read(b buffer.PacketBuffer) {
}

func (packet *ClearChoicePacket) Write(b buffer.PacketBuffer) {
}
