package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type OpenUIPacket struct {
	*wave.DefaultPacket
	Variable1 int32
	Variable2 int32
	Item      string
}

func (packet *OpenUIPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Item = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *OpenUIPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteString(b.Bytes(), packet.Item, b.Index())
}
