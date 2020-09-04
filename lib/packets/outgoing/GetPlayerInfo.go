package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type GetPlayerInfoPacket struct {
	*wave.DefaultPacket
	Index int32
	X     byte
	Y     byte
}

func (packet *GetPlayerInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.X = b.ReadByte(b.Bytes(), b.Index())
	packet.Y = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *GetPlayerInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteByte(b.Bytes(), packet.X, b.Index())
	b.WriteByte(b.Bytes(), packet.Y, b.Index())
}
