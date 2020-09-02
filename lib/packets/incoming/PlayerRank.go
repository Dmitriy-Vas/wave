package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerRankPacket struct {
	*wave.DefaultPacket
	Variable0  int64
	Variable1  byte
	Variable2  bool
	Variable3  bool
	Variable4  string
	Variable5  string
	Variable6  int32
	Variable7  string
	Variable8  int32
	Variable9  int32
	Variable10 string
	Variable11 int32
	Variable12 int64
	Variable13 string
	Variable14 string
	Variable15 int32
	Variable16 int32
	Variable17 int32
	Variable18 string
	Variable19 int32
	Variable20 int32
	Variable21 int64
	Variable22 string
	Variable23 int32
	Variable24 int32
}

func (packet *PlayerRankPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable5 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable6 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable8 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable9 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable10 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable11 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable12 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable13 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable14 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable15 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable16 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable17 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable18 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable19 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable20 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable21 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable22 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable23 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable24 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerRankPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable2, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable3, b.Index())
	b.WriteString(b.Bytes(), packet.Variable4, b.Index())
	b.WriteString(b.Bytes(), packet.Variable5, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6, b.Index())
	b.WriteString(b.Bytes(), packet.Variable7, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable8, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable9, b.Index())
	b.WriteString(b.Bytes(), packet.Variable10, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable11, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable12, b.Index())
	b.WriteString(b.Bytes(), packet.Variable13, b.Index())
	b.WriteString(b.Bytes(), packet.Variable14, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable15, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable16, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable17, b.Index())
	b.WriteString(b.Bytes(), packet.Variable18, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable19, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable20, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable21, b.Index())
	b.WriteString(b.Bytes(), packet.Variable22, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable23, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable24, b.Index())
}
