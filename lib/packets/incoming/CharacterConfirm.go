package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type CharacterConfirmPacket struct {
	*wave.DefaultPacket
}

func (packet *CharacterConfirmPacket) Read(b buffer.PacketBuffer) {
}

func (packet *CharacterConfirmPacket) Write(b buffer.PacketBuffer) {
}
