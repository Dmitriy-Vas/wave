package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PartyQuestInfoPacket struct {
	*wave.DefaultPacket
	Time int32
}

func (packet *PartyQuestInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Time = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PartyQuestInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Time, b.Index())
}
