package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerMasterPacket struct {
	*wave.DefaultPacket
	Item   int32
	Master []int32
}

func (packet *PlayerMasterPacket) Read(b buffer.PacketBuffer) {
	packet.Item = b.ReadInt(b.Bytes(), b.Index())
	packet.Master = make([]int32, 25)
	for i, _ := range packet.Master {
		packet.Master[i] = b.ReadInt(b.Bytes(), b.Index())
	}
}

func (packet *PlayerMasterPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Item, b.Index())
	for _, m := range packet.Master {
		b.WriteInt(b.Bytes(), m, b.Index())
	}
}
