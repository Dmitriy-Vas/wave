package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type StopDefencePacket struct {
	*wave.DefaultPacket
	Variable1 int64
}

func (packet *StopDefencePacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadLong(b.Bytes(), b.Index())

}

func (packet *StopDefencePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable1, b.Index())
}
