package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

// Empty packet
type PartyLeavePacket struct {
	*wave.DefaultPacket
}

func (packet *PartyLeavePacket) Read(b buffer.PacketBuffer) {
}

func (packet *PartyLeavePacket) Write(b buffer.PacketBuffer) {
}
