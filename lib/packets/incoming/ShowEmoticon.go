package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ShowEmoticonPacket struct {
	*wave.DefaultPacket
	IsPin bool
	Index int32
	Num   int32
}

func (packet *ShowEmoticonPacket) Read(b buffer.PacketBuffer) {
	packet.IsPin = b.ReadBool(b.Bytes(), b.Index())
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *ShowEmoticonPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.IsPin, b.Index())
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
}
