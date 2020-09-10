package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdateQuestPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdateQuestPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdateQuestPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdateQuestPacket) SetSend(value bool) {
	d.Send = value
}

type UpdateQuestPacket struct {
	ID        int64
	Send      bool
	Variable1 int32
}

func (packet *UpdateQuestPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	// TODO quest data
}

func (packet *UpdateQuestPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	// TODO quest data
}
