package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UrlPacket struct {
	*wave.DefaultPacket
	PromotionURL string
}

func (packet *UrlPacket) Read(b buffer.PacketBuffer) {
	packet.PromotionURL = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *UrlPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.PromotionURL, b.Index())
}
