package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ShowEmoticonPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 bool
}

func (packet *ShowEmoticonPacket) Read(b buffer.PacketBuffer) {

}

func (packet *ShowEmoticonPacket) Write(b buffer.PacketBuffer) {

}
