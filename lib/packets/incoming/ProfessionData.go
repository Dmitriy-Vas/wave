package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ProfessionDataPacket struct {
	*wave.DefaultPacket
}

func (packet *ProfessionDataPacket) Read(b buffer.PacketBuffer) {
	// TODO profession data
}

func (packet *ProfessionDataPacket) Write(b buffer.PacketBuffer) {
	// TODO profession data
}
