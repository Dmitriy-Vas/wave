package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"time"
)

type GetStatsPacket struct {
	*wave.DefaultPacket
	Variable1 time.Time
	Variable2 time.Time
}

func (packet *GetStatsPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = wrapper.ReadDate(b)
	packet.Variable2 = wrapper.ReadDate(b)

}

func (packet *GetStatsPacket) Write(b buffer.PacketBuffer) {
	wrapper.WriteDate(b, packet.Variable1)
	wrapper.WriteDate(b, packet.Variable2)
}
