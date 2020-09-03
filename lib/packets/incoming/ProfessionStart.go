package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ProfessionStartPacket struct {
	*wave.DefaultPacket
	Variable1    byte
	FishingDiff  int32
	FishingNum   int32
	FishingSpeed int32
}

func (packet *ProfessionStartPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.FishingDiff = b.ReadInt(b.Bytes(), b.Index())
	packet.FishingNum = b.ReadInt(b.Bytes(), b.Index())
	packet.FishingSpeed = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *ProfessionStartPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.FishingDiff, b.Index())
	b.WriteInt(b.Bytes(), packet.FishingNum, b.Index())
	b.WriteInt(b.Bytes(), packet.FishingSpeed, b.Index())
}
