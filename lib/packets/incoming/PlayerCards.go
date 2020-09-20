package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (packet *PlayerCardsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerCardsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerCardsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerCardsPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerCardsPacket struct {
	ID          int64
	Send        bool
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
