package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type BreakPointPacket struct {
	*wave.DefaultPacket
	Num    int32
	Prompt string
}

func (packet *BreakPointPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.Prompt = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *BreakPointPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteString(b.Bytes(), packet.Prompt, b.Index())
}
