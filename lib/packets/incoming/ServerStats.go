package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (packet *ServerStatsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ServerStatsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ServerStatsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ServerStatsPacket) SetSend(value bool) {
	packet.Send = value
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
	for i := range packet.Cards {
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
