package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type TargetPacket struct {
	*wave.DefaultPacket
	Target     int32
	TargetType byte
}

func (packet *TargetPacket) Read(b buffer.PacketBuffer) {
	packet.Target = b.ReadInt(b.Bytes(), b.Index())
	packet.TargetType = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *TargetPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Target, b.Index())
	b.WriteByte(b.Bytes(), packet.TargetType, b.Index())
}
