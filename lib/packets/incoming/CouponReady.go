package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *CouponReadyPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CouponReadyPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CouponReadyPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CouponReadyPacket) SetSend(value bool) {
	d.Send = value
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
