package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type ServerStatsPacket struct {
	*wave.DefaultPacket
	CardNum byte
	Cards   []lib.PromotionGameCardRec
}

func (packet *ServerStatsPacket) Read(b buffer.PacketBuffer) {
	packet.CardNum = b.ReadByte(b.Bytes(), b.Index())
	packet.Cards = make([]lib.PromotionGameCardRec, packet.CardNum)
	for i, _ := range packet.Cards {
		packet.Cards[i].Name = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.Cards[i].Price = b.ReadString(b.Bytes(), b.Index(), 0)
		packet.Cards[i].Bonus = b.ReadInt(b.Bytes(), b.Index())
	}
}

func (packet *ServerStatsPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.CardNum, b.Index())
	for _, card := range packet.Cards {
		b.WriteString(b.Bytes(), card.Name, b.Index())
		b.WriteString(b.Bytes(), card.Price, b.Index())
		b.WriteInt(b.Bytes(), card.Bonus, b.Index())
	}
}
