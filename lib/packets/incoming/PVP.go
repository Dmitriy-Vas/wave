package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type PVPPacket struct {
	*wave.DefaultPacket
	Num  int64
	Num2 int64
}

func (packet *PVPPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadLong(b.Bytes(), b.Index())
	packet.Num2 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *PVPPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Num, b.Index())
	b.WriteLong(b.Bytes(), packet.Num2, b.Index())
}
