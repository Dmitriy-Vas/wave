package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type SetAccessPacket struct {
	*wave.DefaultPacket
	Name   string
	Access int64
}

func (packet *SetAccessPacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Access = b.ReadLong(b.Bytes(), b.Index())

}

func (packet *SetAccessPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
	b.WriteLong(b.Bytes(), packet.Access, b.Index())
}
