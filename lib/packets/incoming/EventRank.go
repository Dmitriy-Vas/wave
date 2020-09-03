package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

type EventRankPacket struct {
	*wave.DefaultPacket
	Events []string
	Ranks  []lib.EventsRankRec
}

func (packet *EventRankPacket) Read(b buffer.PacketBuffer) {
	packet.Events = make([]string, 3) // TODO int to const
	for i, _ := range packet.Events {
		packet.Events[i] = b.ReadString(b.Bytes(), b.Index(), 0)
	}
	packet.Ranks = make([]lib.EventsRankRec, 10) // TODO int to const
	for i, _ := range packet.Ranks {
		packet.Ranks[i] = lib.EventsRankRec{
			Name:    b.ReadString(b.Bytes(), b.Index(), 0),
			Classes: b.ReadInt(b.Bytes(), b.Index()),
			Score:   b.ReadLong(b.Bytes(), b.Index()),
		}
	}
}

func (packet *EventRankPacket) Write(b buffer.PacketBuffer) {
	for _, event := range packet.Events {
		b.WriteString(b.Bytes(), event, b.Index())
	}
	for _, rank := range packet.Ranks {
		b.WriteString(b.Bytes(), rank.Name, b.Index())
		b.WriteInt(b.Bytes(), rank.Classes, b.Index())
		b.WriteLong(b.Bytes(), rank.Score, b.Index())
	}
}
