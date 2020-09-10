package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ProfessionDataPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ProfessionDataPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ProfessionDataPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ProfessionDataPacket) SetSend(value bool) {
	d.Send = value
}

type ProfessionDataPacket struct {
	ID   int64
	Send bool
}

func (packet *ProfessionDataPacket) Read(b buffer.PacketBuffer) {
	// TODO profession data
}

func (packet *ProfessionDataPacket) Write(b buffer.PacketBuffer) {
	// TODO profession data
}
