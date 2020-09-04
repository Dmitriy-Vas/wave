package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type AdminTaskPacket struct {
	*wave.DefaultPacket
	TaskNum   byte
	Variable2 int32
	Variable3 string
}

func (packet *AdminTaskPacket) Read(b buffer.PacketBuffer) {
	packet.TaskNum = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *AdminTaskPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.TaskNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteString(b.Bytes(), packet.Variable3, b.Index())
}
