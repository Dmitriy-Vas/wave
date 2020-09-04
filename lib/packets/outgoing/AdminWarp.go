package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type AdminWarpPacket struct {
	*wave.DefaultPacket
	X int32
	Y int32
}

func (packet *AdminWarpPacket) Read(b buffer.PacketBuffer) {
	packet.X = b.ReadInt(b.Bytes(), b.Index())
	packet.Y = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *AdminWarpPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Y, b.Index())
}
