package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (d *EventRankPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *EventRankPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *EventRankPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *EventRankPacket) SetSend(value bool) {
	d.Send = value
}

type EventRankPacket struct {
	ID     int64
	Send   bool
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
