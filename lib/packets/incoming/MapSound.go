package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type MapSoundPacket struct {
	*wave.DefaultPacket
	Num   int32
	Sound string
}

func (packet *MapSoundPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.Sound = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *MapSoundPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteString(b.Bytes(), packet.Sound, b.Index())
}
