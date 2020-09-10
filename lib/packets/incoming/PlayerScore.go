package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerScorePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerScorePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerScorePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerScorePacket) SetSend(value bool) {
	d.Send = value
}

type PlayerScorePacket struct {
	ID          int64
	Send        bool
	SoccerScore int32
}

func (packet *PlayerScorePacket) Read(b buffer.PacketBuffer) {
	packet.SoccerScore = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerScorePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.SoccerScore, b.Index())
}
