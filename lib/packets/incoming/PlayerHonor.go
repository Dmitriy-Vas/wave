package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerHonorPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 int32
	Variable2 int32
}

func (packet *PlayerHonorPacket) Read(b buffer.PacketBuffer) {

}

func (packet *PlayerHonorPacket) Write(b buffer.PacketBuffer) {

}
