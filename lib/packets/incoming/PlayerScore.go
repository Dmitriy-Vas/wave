package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PlayerScorePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerScorePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerScorePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerScorePacket) SetSend(value bool) {
	packet.Send = value
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
