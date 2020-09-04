package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type WarpMapPacket struct {
	*wave.DefaultPacket
	MapIndex int32
}

func (packet *WarpMapPacket) Read(b buffer.PacketBuffer) {
	packet.MapIndex = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *WarpMapPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MapIndex, b.Index())
}
