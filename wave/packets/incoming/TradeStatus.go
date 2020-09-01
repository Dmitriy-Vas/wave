package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type TradeStatusPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 string
	Variable2 byte
}

func (packet *TradeStatusPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *TradeStatusPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteString(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
}
