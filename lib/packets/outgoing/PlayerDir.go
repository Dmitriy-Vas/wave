package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PlayerDirPacket struct {
	*wave.DefaultPacket
	MyIndex int32
	Dir     byte
}

func (packet *PlayerDirPacket) Read(b buffer.PacketBuffer) {
	packet.MyIndex = b.ReadInt(b.Bytes(), b.Index())
	packet.Dir = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *PlayerDirPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MyIndex, b.Index())
	b.WriteByte(b.Bytes(), packet.Dir, b.Index())
}
