package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type FlashPacket struct {
	*wave.DefaultPacket
	Variable1  byte
	Variable2  int32
	Variable3  byte
	Expression bool
	Num        int32
}

func (packet *FlashPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadByte(b.Bytes(), b.Index())
	packet.Expression = b.ReadBool(b.Bytes(), b.Index())
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *FlashPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable3, b.Index())
	b.WriteBool(b.Bytes(), packet.Expression, b.Index())
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
}
