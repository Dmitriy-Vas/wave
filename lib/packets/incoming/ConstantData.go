package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ConstantDataPacket struct {
	*wave.DefaultPacket
	MaxPlayers int32
	MaxLevels  int32
}

func (packet *ConstantDataPacket) Read(b buffer.PacketBuffer) {
	packet.MaxPlayers = b.ReadInt(b.Bytes(), b.Index())
	packet.MaxLevels = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *ConstantDataPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MaxPlayers, b.Index())
	b.WriteInt(b.Bytes(), packet.MaxLevels, b.Index())
}
