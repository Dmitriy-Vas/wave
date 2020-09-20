package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ProfessionDataPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ProfessionDataPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ProfessionDataPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ProfessionDataPacket) SetSend(value bool) {
	packet.Send = value
}

type ProfessionDataPacket struct {
	ID   int64
	Send bool
}

func (packet *ProfessionDataPacket) Read(_ buffer.PacketBuffer) {
	// TODO profession data
}

func (packet *ProfessionDataPacket) Write(_ buffer.PacketBuffer) {
	// TODO profession data
}
