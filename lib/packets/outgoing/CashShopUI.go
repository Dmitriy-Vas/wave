package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *CashShopUIPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *CashShopUIPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *CashShopUIPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *CashShopUIPacket) SetSend(value bool) {
	packet.Send = value
}

type CashShopUIPacket struct {
	ID     int64
	Send   bool
	Action bool
}

func (packet *CashShopUIPacket) Read(b buffer.PacketBuffer) {
	packet.Action = b.ReadBool(b.Bytes(), b.Index())

}

func (packet *CashShopUIPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.Action, b.Index())
}
