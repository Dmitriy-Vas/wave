package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type CheckMapPacket struct {
	*wave.DefaultPacket
	MapNum      int32
	MapRevision int32
	X           int32
	Y           int32
}

func (packet *CheckMapPacket) Read(b buffer.PacketBuffer) {
	packet.MapNum = b.ReadInt(b.Bytes(), b.Index())
	packet.MapRevision = b.ReadInt(b.Bytes(), b.Index())
	packet.X = b.ReadInt(b.Bytes(), b.Index())
	packet.Y = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *CheckMapPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MapNum, b.Index())
	b.WriteInt(b.Bytes(), packet.MapRevision, b.Index())
	b.WriteInt(b.Bytes(), packet.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Y, b.Index())
}
