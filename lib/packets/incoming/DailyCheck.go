package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"time"
)

type DailyCheckPacket struct {
	*wave.DefaultPacket
	CheckPlayerDate time.Time
	CheckCount      byte
	CheckDate       time.Time
}

func (packet *DailyCheckPacket) Read(b buffer.PacketBuffer) {
	packet.CheckPlayerDate = wrapper.ReadDate(b)
	packet.CheckCount = b.ReadByte(b.Bytes(), b.Index())
	packet.CheckDate = wrapper.ReadDate(b)

}

func (packet *DailyCheckPacket) Write(b buffer.PacketBuffer) {
	wrapper.WriteDate(b, packet.CheckPlayerDate)
	b.WriteByte(b.Bytes(), packet.CheckCount, b.Index())
	wrapper.WriteDate(b, packet.CheckDate)
}
