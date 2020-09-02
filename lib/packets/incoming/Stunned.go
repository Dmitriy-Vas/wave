package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type StunnedPacket struct {
	*wave.DefaultPacket
	StunDuration int32
	StopDuration bool
}

func (packet *StunnedPacket) Read(b buffer.PacketBuffer) {
	packet.StunDuration = b.ReadInt(b.Bytes(), b.Index())
	packet.StopDuration = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *StunnedPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.StunDuration, b.Index())
	b.WriteBool(b.Bytes(), packet.StopDuration, b.Index())
}
