package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (d *ServerStatsPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ServerStatsPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ServerStatsPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ServerStatsPacket) SetSend(value bool) {
	d.Send = value
}

type ServerStatsPacket struct {
	ID      int64
	Send    bool
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
