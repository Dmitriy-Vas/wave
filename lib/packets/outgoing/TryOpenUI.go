package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type TryOpenUIPacket struct {
	*wave.DefaultPacket
	UiIndex int32
}

func (packet *TryOpenUIPacket) Read(b buffer.PacketBuffer) {
	packet.UiIndex = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *TryOpenUIPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.UiIndex, b.Index())
}
