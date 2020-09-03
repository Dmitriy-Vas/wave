package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerConnectedPacket struct {
	*wave.DefaultPacket
	Amount    int32
	Variable2 bool
	Variable3 []string
}

func (packet *PlayerConnectedPacket) Read(b buffer.PacketBuffer) {
	packet.Amount = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadBool(b.Bytes(), b.Index())
	if packet.Variable2 {
		packet.Variable3 = make([]string, packet.Amount)
		for i, _ := range packet.Variable3 {
			packet.Variable3[i] = b.ReadString(b.Bytes(), b.Index(), 0)
		}
	}
}

func (packet *PlayerConnectedPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Amount, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable2, b.Index())
	if packet.Variable2 {
		for _, text := range packet.Variable3 {
			b.WriteString(b.Bytes(), text, b.Index())
		}
	}
}
