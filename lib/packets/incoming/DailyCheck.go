package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"time"
)

type DailyCheckPacket struct {
	*wave.DefaultPacket
	CheckPlayerDate time.Time
	CheckCount      byte
	CheckDate       time.Time
}

func (packet *DailyCheckPacket) Read(b buffer.PacketBuffer) {

}

func (packet *DailyCheckPacket) Write(b buffer.PacketBuffer) {

}
