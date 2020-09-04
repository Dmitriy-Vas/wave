package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UseItemPacket struct {
	*wave.DefaultPacket
	InvNum int32
}

func (packet *UseItemPacket) Read(b buffer.PacketBuffer) {
	packet.InvNum = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *UseItemPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.InvNum, b.Index())
}
