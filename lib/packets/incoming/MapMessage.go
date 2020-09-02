package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type MapMessagePacket struct {
	*wave.DefaultPacket
	Message string
}

func (packet *MapMessagePacket) Read(b buffer.PacketBuffer) {
	packet.Message = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *MapMessagePacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Message, b.Index())
}
