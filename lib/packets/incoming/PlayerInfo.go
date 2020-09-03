package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerInfoPacket struct {
	*wave.DefaultPacket
	Index     int32
	Variable2 []int32
	Variable3 []string
	ShowStats bool
}

func (packet *PlayerInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = make([]int32, 7) // TODO int to const
	packet.Variable3 = make([]string, 7)
	for i, _ := range packet.Variable2 {
		packet.Variable2[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Variable3[i] = b.ReadString(b.Bytes(), b.Index(), 0)
	}
	packet.ShowStats = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *PlayerInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	for i, _ := range packet.Variable2 {
		b.WriteInt(b.Bytes(), packet.Variable2[i], b.Index())
		b.WriteString(b.Bytes(), packet.Variable3[i], b.Index())
	}
	b.WriteBool(b.Bytes(), packet.ShowStats, b.Index())
}
