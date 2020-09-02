package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerCraftingPacket struct {
	*wave.DefaultPacket
	Num     int32
	WithLog bool
	Num2    int32
	Amount  int32
}

func (packet *PlayerCraftingPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.WithLog = b.ReadBool(b.Bytes(), b.Index())
	packet.Num2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Amount = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerCraftingPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteBool(b.Bytes(), packet.WithLog, b.Index())
	b.WriteInt(b.Bytes(), packet.Num2, b.Index())
	b.WriteInt(b.Bytes(), packet.Amount, b.Index())
}
