package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/buffer/objects"
)

type MapNPCDataPacket struct {
	*wave.DefaultPacket
	Num     int32
	Pos     objects.Vector2
	Dir     byte
	Vital   int32
	HPSetTo int32
}

func (packet *MapNPCDataPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.Pos = b.ReadVector2(b.Bytes(), b.Index())
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
	packet.Vital = b.ReadInt(b.Bytes(), b.Index())
	packet.HPSetTo = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *MapNPCDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteVector2(b.Bytes(), packet.Pos, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
	b.WriteInt(b.Bytes(), packet.Vital, b.Index())
	b.WriteInt(b.Bytes(), packet.HPSetTo, b.Index())
}
