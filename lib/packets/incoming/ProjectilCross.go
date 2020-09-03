package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ProjectilCrossPacket struct {
	*wave.DefaultPacket
	PlayerNum      int64
	ProjectilNum   int64
	ProjectilIndex int64
	Combo          int64
	Direction      int64
}

func (packet *ProjectilCrossPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadLong(b.Bytes(), b.Index())
	packet.ProjectilNum = b.ReadLong(b.Bytes(), b.Index())
	packet.ProjectilIndex = b.ReadLong(b.Bytes(), b.Index())
	packet.Combo = b.ReadLong(b.Bytes(), b.Index())
	packet.Direction = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *ProjectilCrossPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteLong(b.Bytes(), packet.ProjectilNum, b.Index())
	b.WriteLong(b.Bytes(), packet.ProjectilIndex, b.Index())
	b.WriteLong(b.Bytes(), packet.Combo, b.Index())
	b.WriteLong(b.Bytes(), packet.Direction, b.Index())
}
