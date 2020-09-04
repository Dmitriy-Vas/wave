package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UseMegaphonePacket struct {
	*wave.DefaultPacket
	Type    byte
	Message string
	InvItem byte
}

func (packet *UseMegaphonePacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Message = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.InvItem = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *UseMegaphonePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteString(b.Bytes(), packet.Message, b.Index())
	b.WriteByte(b.Bytes(), packet.InvItem, b.Index())
}
