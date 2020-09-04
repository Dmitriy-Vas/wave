package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type TradeRequestPacket struct {
	*wave.DefaultPacket
	CurX int32
}

func (packet *TradeRequestPacket) Read(b buffer.PacketBuffer) {
	packet.CurX = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *TradeRequestPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.CurX, b.Index())
}
