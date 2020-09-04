package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type RequestDataPacket struct {
	*wave.DefaultPacket
	DataType int32
}

func (packet *RequestDataPacket) Read(b buffer.PacketBuffer) {
	packet.DataType = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *RequestDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.DataType, b.Index())
}
