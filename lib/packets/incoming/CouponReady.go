package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type CouponReadyPacket struct {
	*wave.DefaultPacket
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
