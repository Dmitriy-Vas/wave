package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type EventRankPacket struct {
	*wave.DefaultPacket
}

func (packet *EventRankPacket) Read(b buffer.PacketBuffer) {

}

func (packet *EventRankPacket) Write(b buffer.PacketBuffer) {

}
