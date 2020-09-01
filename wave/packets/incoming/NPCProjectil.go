package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type NPCProjectilPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 int32
	Variable2 byte
	Variable3 int32
	Variable4 int32
	Variable5 byte
	Variable6 int32
	Variable7 int32
	Variable8 byte
	Variable9 byte
}

func (packet *NPCProjectilPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable8 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable9 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *NPCProjectilPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable4, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable5, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable7, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable8, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable9, b.Index())
}
