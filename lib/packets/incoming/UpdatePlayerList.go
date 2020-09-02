package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdatePlayerListPacket struct {
	*wave.DefaultPacket
	Index     int32
	Variable1 int32
	Type      int32
}

func (packet *UpdatePlayerListPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Type = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *UpdatePlayerListPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Type, b.Index())
}
