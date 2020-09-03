package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type AnnouncerPacket struct {
	*wave.DefaultPacket
}

func (packet *AnnouncerPacket) Read(b buffer.PacketBuffer) {

}

func (packet *AnnouncerPacket) Write(b buffer.PacketBuffer) {

}
