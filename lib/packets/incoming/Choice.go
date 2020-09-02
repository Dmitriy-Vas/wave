package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/buffer/objects"
)

type ChoicePacket struct {
	*wave.DefaultPacket
	Variable0  int64
	Variable1  bool
	Variable2  int32
	Variable3  objects.Vector2
	Variable4  bool
	Variable5  string
	Variable6  byte
	Variable7  bool
	Variable8  objects.Vector2
	Variable9  int32
	Variable10 string
}

func (packet *ChoicePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadVector2(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable6 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable8 = b.ReadVector2(b.Bytes(), b.Index())
	packet.Variable9 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable10 = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *ChoicePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteVector2(b.Bytes(), packet.Variable3, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable4, b.Index())
	b.WriteString(b.Bytes(), packet.Variable5, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable6, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable7, b.Index())
	b.WriteVector2(b.Bytes(), packet.Variable8, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable9, b.Index())
	b.WriteString(b.Bytes(), packet.Variable10, b.Index())
}
