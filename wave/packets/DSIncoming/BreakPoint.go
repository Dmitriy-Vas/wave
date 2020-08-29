package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type BreakPointPacket struct {
	*packets.DefaultPacket
	Num    int32
	Prompt string
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &BreakPointPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *BreakPointPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.Prompt = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *BreakPointPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteString(b.Bytes(), packet.Prompt, b.Index())
}
