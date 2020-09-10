package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *TradeRequestPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *TradeRequestPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *TradeRequestPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *TradeRequestPacket) SetSend(value bool) {
	d.Send = value
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
