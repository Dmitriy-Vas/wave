package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type PlayerDataPacket struct {
	*packets.DefaultPacket
	Variable0  int64
	Variable1  int32
	Variable2  bool
	Variable3  string
	Variable4  int32
	Variable5  byte
	Variable6  int32
	Variable7  int32
	Variable8  int32
	Variable9  int32
	Variable10 int32
	Variable11 int32
	Variable12 byte
	Variable13 byte
	Variable14 byte
	Variable15 int32
	Variable16 int32
	Variable17 string
	Variable18
	Variable19 int64
	Variable20 int64
	Variable21 int64
	Variable22 int64
	Variable23 byte
	Variable24 byte
	Variable25 byte
	Variable26 byte
	Variable27 int32
	Variable28 int32
	Variable29 int32
	Variable30 int32
	Variable31 int32
	Variable32 int32
	Variable33 int32
	Variable34 int32
	Variable35
	Variable36 int32
	Variable37 byte
	Variable38 byte
	Variable39 int32
	Variable40 int32
	Variable41 bool
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &PlayerDataPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *PlayerDataPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable4 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable7 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable8 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable9 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable10 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable11 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable12 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable13 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable14 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable15 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable16 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable17 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable18 = b.Read(b.Bytes(), b.Index())
	packet.Variable19 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable20 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable21 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable22 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable23 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable24 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable25 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable26 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable27 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable28 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable29 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable30 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable31 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable32 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable33 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable34 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable35 = b.Read(b.Bytes(), b.Index())
	packet.Variable36 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable37 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable38 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable39 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable40 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable41 = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *PlayerDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable2, b.Index())
	b.WriteString(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable4, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable5, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable6, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable7, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable8, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable9, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable10, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable11, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable12, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable13, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable14, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable15, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable16, b.Index())
	b.WriteString(b.Bytes(), packet.Variable17, b.Index())
	b.Write(b.Bytes(), packet.Variable18, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable19, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable20, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable21, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable22, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable23, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable24, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable25, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable26, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable27, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable28, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable29, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable30, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable31, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable32, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable33, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable34, b.Index())
	b.Write(b.Bytes(), packet.Variable35, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable36, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable37, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable38, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable39, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable40, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable41, b.Index())
}
