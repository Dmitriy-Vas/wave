package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *CashShopUIPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CashShopUIPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CashShopUIPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CashShopUIPacket) SetSend(value bool) {
	d.Send = value
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
