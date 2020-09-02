package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SpellAnimPacket struct {
	*wave.DefaultPacket
	PlayerNum int32
	AnimNum   int32
}

func (packet *SpellAnimPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.AnimNum = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *SpellAnimPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.AnimNum, b.Index())
}
