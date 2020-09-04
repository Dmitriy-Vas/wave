package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type DemoValuePacket struct {
	*wave.DefaultPacket
	ValType byte
	ValOne  int32
	ValTwo  int32
}

func (packet *DemoValuePacket) Read(b buffer.PacketBuffer) {
	packet.ValType = b.ReadByte(b.Bytes(), b.Index())
	packet.ValOne = b.ReadInt(b.Bytes(), b.Index())
	packet.ValTwo = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *DemoValuePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.ValType, b.Index())
	b.WriteInt(b.Bytes(), packet.ValOne, b.Index())
	b.WriteInt(b.Bytes(), packet.ValTwo, b.Index())
}
