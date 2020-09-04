package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type WhisperPacket struct {
	*wave.DefaultPacket
	Name string
	Text string
}

func (packet *WhisperPacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *WhisperPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
	b.WriteString(b.Bytes(), packet.Text, b.Index())
}
