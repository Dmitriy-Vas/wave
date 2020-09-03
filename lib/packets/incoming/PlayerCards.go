package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type PlayerCardsPacket struct {
	*wave.DefaultPacket
	PlayerNum   int32
	CardNum     int32
	PlayerCards []lib.PlayerCardRec
	CurrentCard int32
}

func (packet *PlayerCardsPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.CardNum = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayerCards = make([]lib.PlayerCardRec, 255) // TODO move to constants
	for i := range packet.PlayerCards {
		packet.PlayerCards[i] = lib.PlayerCardRec{
			Level: b.ReadInt(b.Bytes(), b.Index()),
			Exp:   b.ReadInt(b.Bytes(), b.Index()),
		}
	}
	packet.CurrentCard = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerCardsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.CardNum, b.Index())
	for _, card := range packet.PlayerCards {
		b.WriteInt(b.Bytes(), card.Level, b.Index())
		b.WriteInt(b.Bytes(), card.Exp, b.Index())
	}
	b.WriteInt(b.Bytes(), packet.CurrentCard, b.Index())
}
