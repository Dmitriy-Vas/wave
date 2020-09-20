package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PartyQuestInfoPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PartyQuestInfoPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PartyQuestInfoPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PartyQuestInfoPacket) SetSend(value bool) {
	packet.Send = value
}

type PartyQuestInfoPacket struct {
	ID   int64
	Send bool
	Time int32
}

func (packet *PartyQuestInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Time = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PartyQuestInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Time, b.Index())
}
