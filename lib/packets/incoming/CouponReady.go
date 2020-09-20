package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *CouponReadyPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *CouponReadyPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *CouponReadyPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *CouponReadyPacket) SetSend(value bool) {
	packet.Send = value
}

type CouponReadyPacket struct {
	ID          int64
	Send        bool
	Coupon      byte
	CouponValue int32
}

func (packet *CouponReadyPacket) Read(b buffer.PacketBuffer) {
	packet.Coupon = b.ReadByte(b.Bytes(), b.Index())
	packet.CouponValue = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *CouponReadyPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Coupon, b.Index())
	b.WriteInt(b.Bytes(), packet.CouponValue, b.Index())
}
