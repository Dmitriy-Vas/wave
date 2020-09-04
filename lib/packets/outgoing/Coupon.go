package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type CouponPacket struct {
	*wave.DefaultPacket
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
