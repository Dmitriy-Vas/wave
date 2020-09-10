package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *CouponPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CouponPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CouponPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CouponPacket) SetSend(value bool) {
	d.Send = value
}

type CouponPacket struct {
	ID            int64
	Send          bool
	Code          string
	WaitingCoupon bool
}

func (packet *CouponPacket) Read(b buffer.PacketBuffer) {
	packet.Code = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.WaitingCoupon = b.ReadBool(b.Bytes(), b.Index())

}

func (packet *CouponPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Code, b.Index())
	b.WriteBool(b.Bytes(), packet.WaitingCoupon, b.Index())
}
