package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type BanListPacket struct {
	*wave.DefaultPacket
}

func (packet *BanListPacket) Read(b buffer.PacketBuffer) {

}

func (packet *BanListPacket) Write(b buffer.PacketBuffer) {
}
