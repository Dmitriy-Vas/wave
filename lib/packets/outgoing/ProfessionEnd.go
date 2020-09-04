package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ProfessionEndPacket struct {
	*wave.DefaultPacket
	FishingCount int32
	FishingPos   int32
	FishingFail  byte
	FishingTime  int32
}

func (packet *ProfessionEndPacket) Read(b buffer.PacketBuffer) {
	packet.FishingCount = b.ReadInt(b.Bytes(), b.Index())
	packet.FishingPos = b.ReadInt(b.Bytes(), b.Index())
	packet.FishingFail = b.ReadByte(b.Bytes(), b.Index())
	packet.FishingTime = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *ProfessionEndPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.FishingCount, b.Index())
	b.WriteInt(b.Bytes(), packet.FishingPos, b.Index())
	b.WriteByte(b.Bytes(), packet.FishingFail, b.Index())
	b.WriteInt(b.Bytes(), packet.FishingTime, b.Index())
}
