package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ProjectilPacket struct {
	*wave.DefaultPacket
	ProjectilNum int32
	PlayerNum    int32
	Variable3    int32
	Combo        int32
	Dir          byte
}

func (packet *ProjectilPacket) Read(b buffer.PacketBuffer) {
	packet.ProjectilNum = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
	packet.Combo = b.ReadInt(b.Bytes(), b.Index())
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ProjectilPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ProjectilNum, b.Index())
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Combo, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
}
