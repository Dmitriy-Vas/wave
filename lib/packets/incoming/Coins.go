package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type CoinsPacket struct {
	*wave.DefaultPacket
	PlayerCoins int32
}

func (packet *CoinsPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerCoins = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *CoinsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerCoins, b.Index())
}
