package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/objects"
)

type MapPacket struct {
	*wave.DefaultPacket
	Variable1  int32
	Variable2  string
	Variable3  string
	Variable4  byte
	Variable5  byte
	Variable6  int32
	Variable7  int32
	Variable8  int32
	Variable9  int32
	Variable10 int32
	Variable11 int32
	Variable12 int32
	Variable13 int32
	Variable14 int32
	Variable15 int32
	Variable16 int32
	Variable17 int32
	Variable18 byte
	Variable19 int32
	Variable20 int32
	Variable21 int32
	Variable22 int32
	Variable23 int32
	Variable24 byte
	Variable25 string
	Variable26 int32
	Variable27 int32
	Variable28 int32
	Variable29 objects.Vector2
	Variable30 objects.Vector2
	Variable31 objects.Vector2
	Variable32 objects.Vector2
	Variable33 objects.Vector2
	Variable34 int32
	Variable35 objects.Vector2
	Variable36 int64
	Variable37 int64
	Variable38 int64
	Variable39 int64
	Variable40 string
	Variable41 int64
	Variable42 int64
	Variable43 int64
	Variable44 int32
	Variable45 string
	Variable46 int64
	Variable47 []byte
}

func (packet *MapPacket) Read(b buffer.PacketBuffer) {
	// packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable2 = b.ReadString(b.Bytes(), b.Index(), 0)
	// packet.Variable3 = b.ReadString(b.Bytes(), b.Index(), 0)
	// packet.Variable4 = b.ReadByte(b.Bytes(), b.Index())
	// packet.Variable5 = b.ReadByte(b.Bytes(), b.Index())
	// packet.Variable6 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable7 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable8 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable9 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable10 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable11 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable12 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable13 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable14 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable15 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable16 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable17 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable18 = b.ReadByte(b.Bytes(), b.Index())
	// packet.Variable19 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable20 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable21 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable22 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable23 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable24 = b.ReadByte(b.Bytes(), b.Index())
	// packet.Variable25 = b.ReadString(b.Bytes(), b.Index(), 0)
	// packet.Variable26 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable27 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable28 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable29 = objects.Vector2{X: b.ReadInt(b.Bytes(), b.Index()), Y: b.ReadInt(b.Bytes(), b.Index())}
	// packet.Variable30 = objects.Vector2{X: b.ReadInt(b.Bytes(), b.Index()), Y: b.ReadInt(b.Bytes(), b.Index())}
	// packet.Variable31 = objects.Vector2{X: b.ReadInt(b.Bytes(), b.Index()), Y: b.ReadInt(b.Bytes(), b.Index())}
	// packet.Variable32 = objects.Vector2{X: b.ReadInt(b.Bytes(), b.Index()), Y: b.ReadInt(b.Bytes(), b.Index())}
	// packet.Variable33 = objects.Vector2{X: b.ReadInt(b.Bytes(), b.Index()), Y: b.ReadInt(b.Bytes(), b.Index())}
	// packet.Variable34 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable35 = objects.Vector2{X: b.ReadInt(b.Bytes(), b.Index()), Y: b.ReadInt(b.Bytes(), b.Index())}
	// packet.Variable36 = b.ReadLong(b.Bytes(), b.Index())
	// packet.Variable37 = b.ReadLong(b.Bytes(), b.Index())
	// packet.Variable38 = b.ReadLong(b.Bytes(), b.Index())
	// packet.Variable39 = b.ReadLong(b.Bytes(), b.Index())
	// packet.Variable40 = b.ReadString(b.Bytes(), b.Index(), 0)
	// packet.Variable41 = b.ReadLong(b.Bytes(), b.Index())
	// packet.Variable42 = b.ReadLong(b.Bytes(), b.Index())
	// packet.Variable43 = b.ReadLong(b.Bytes(), b.Index())
	// packet.Variable44 = b.ReadInt(b.Bytes(), b.Index())
	// packet.Variable45 = b.ReadString(b.Bytes(), b.Index(), 0)
	// packet.Variable46 = b.ReadLong(b.Bytes(), b.Index())
	// packet.Variable47 = b.ReadBytes(b.Bytes(), b.Index(), 0)

}

func (packet *MapPacket) Write(b buffer.PacketBuffer) {
	// b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	// b.WriteString(b.Bytes(), packet.Variable2, b.Index())
	// b.WriteString(b.Bytes(), packet.Variable3, b.Index())
	// b.WriteByte(b.Bytes(), packet.Variable4, b.Index())
	// b.WriteByte(b.Bytes(), packet.Variable5, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable6, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable7, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable8, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable9, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable10, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable11, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable12, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable13, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable14, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable15, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable16, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable17, b.Index())
	// b.WriteByte(b.Bytes(), packet.Variable18, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable19, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable20, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable21, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable22, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable23, b.Index())
	// b.WriteByte(b.Bytes(), packet.Variable24, b.Index())
	// b.WriteString(b.Bytes(), packet.Variable25, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable26, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable27, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable28, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable29.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable29.Y, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable30.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable30.Y, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable31.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable31.Y, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable32.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable32.Y, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable33.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable33.Y, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable34, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable35.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable35.Y, b.Index())
	// b.WriteLong(b.Bytes(), packet.Variable36, b.Index())
	// b.WriteLong(b.Bytes(), packet.Variable37, b.Index())
	// b.WriteLong(b.Bytes(), packet.Variable38, b.Index())
	// b.WriteLong(b.Bytes(), packet.Variable39, b.Index())
	// b.WriteString(b.Bytes(), packet.Variable40, b.Index())
	// b.WriteLong(b.Bytes(), packet.Variable41, b.Index())
	// b.WriteLong(b.Bytes(), packet.Variable42, b.Index())
	// b.WriteLong(b.Bytes(), packet.Variable43, b.Index())
	// b.WriteInt(b.Bytes(), packet.Variable44, b.Index())
	// b.WriteString(b.Bytes(), packet.Variable45, b.Index())
	// b.WriteLong(b.Bytes(), packet.Variable46, b.Index())
	// b.WriteBytes(b.Bytes(), packet.Variable47, b.Index())
}
