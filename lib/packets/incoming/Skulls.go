package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SkullsPacket struct {
	*wave.DefaultPacket
	PlayerCalaveras []int64
	PickUp          bool
	Variable3       int32
}

func (packet *SkullsPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerCalaveras = make([]int64, 2)
	for i := range packet.PlayerCalaveras {
		packet.PlayerCalaveras[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	packet.PickUp = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *SkullsPacket) Write(b buffer.PacketBuffer) {
	for _, pc := range packet.PlayerCalaveras {
		b.WriteLong(b.Bytes(), pc, b.Index())
	}
	b.WriteBool(b.Bytes(), packet.PickUp, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
}
