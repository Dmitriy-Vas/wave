package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type FavAwardsPacket struct {
	*wave.DefaultPacket
	Action   bool
	AwardNum int32
}

func (packet *FavAwardsPacket) Read(b buffer.PacketBuffer) {
	packet.Action = b.ReadBool(b.Bytes(), b.Index())
	packet.AwardNum = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *FavAwardsPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.Action, b.Index())
	b.WriteInt(b.Bytes(), packet.AwardNum, b.Index())
}
