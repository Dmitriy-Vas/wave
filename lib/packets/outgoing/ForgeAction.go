package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ForgeActionPacket struct {
	*wave.DefaultPacket
	ForgeTimer byte
}

func (packet *ForgeActionPacket) Read(b buffer.PacketBuffer) {
	packet.ForgeTimer = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *ForgeActionPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.ForgeTimer, b.Index())
}
