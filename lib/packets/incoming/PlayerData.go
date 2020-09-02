package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerDataPacket struct {
	*wave.DefaultPacket
}

func (packet *PlayerDataPacket) Read(b buffer.PacketBuffer) {

}

func (packet *PlayerDataPacket) Write(b buffer.PacketBuffer) {

}
