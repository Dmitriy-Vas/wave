package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *TradeRequestPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *TradeRequestPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *TradeRequestPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *TradeRequestPacket) SetSend(value bool) {
	packet.Send = value
}

type TradeRequestPacket struct {
	ID   int64
	Send bool
	CurX int32
}

func (packet *TradeRequestPacket) Read(b buffer.PacketBuffer) {
	packet.CurX = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *TradeRequestPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.CurX, b.Index())
}
