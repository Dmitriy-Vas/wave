package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type BreakPointPacket struct {
	*wave.DefaultPacket
	Line int32
}

func (packet *BreakPointPacket) Read(b buffer.PacketBuffer) {
	packet.Line = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *BreakPointPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Line, b.Index())
}
