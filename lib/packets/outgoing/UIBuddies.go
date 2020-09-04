package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UIBuddiesPacket struct {
	*wave.DefaultPacket
	Action    byte
	ValueStr  string
	ValueByte byte
}

func (packet *UIBuddiesPacket) Read(b buffer.PacketBuffer) {
	packet.Action = b.ReadByte(b.Bytes(), b.Index())
	switch packet.Action {
	case 1:
		packet.ValueStr = b.ReadString(b.Bytes(), b.Index(), 0)
	case 3:
		packet.ValueByte = b.ReadByte(b.Bytes(), b.Index())
	}

}

func (packet *UIBuddiesPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Action, b.Index())
	switch packet.Action {
	case 1:
		b.WriteString(b.Bytes(), packet.ValueStr, b.Index())
	case 3:
		b.WriteByte(b.Bytes(), packet.ValueByte, b.Index())
	}
}
