package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type CharDataPacket struct {
	*wave.DefaultPacket
	Variable0  int64
	Variable1  int32
	Variable2  string
	Variable3  int32
	Variable4  byte
	Variable5  byte
	Variable6  int32
	Variable7  int32
	Variable8  string
	Variable9  int32
	Variable10 int32
	Variable11 bool
	Variable12 bool
}

func (packet *CharDataPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable8 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable9 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable10 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable11 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable12 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *CharDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteString(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable4, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable5, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable7, b.Index())
	b.WriteString(b.Bytes(), packet.Variable8, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable9, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable10, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable11, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable12, b.Index())
}
