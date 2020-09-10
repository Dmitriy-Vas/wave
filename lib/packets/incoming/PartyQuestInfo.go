package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PartyQuestInfoPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PartyQuestInfoPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PartyQuestInfoPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PartyQuestInfoPacket) SetSend(value bool) {
	d.Send = value
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
