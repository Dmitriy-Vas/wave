package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type TradeStatusPacket struct {
	*wave.DefaultPacket
	Text      string
	Variable2 byte
}

func (packet *TradeStatusPacket) Read(b buffer.PacketBuffer) {

	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *TradeStatusPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
}
