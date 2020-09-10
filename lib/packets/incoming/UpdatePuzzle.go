package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdatePuzzlePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdatePuzzlePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdatePuzzlePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdatePuzzlePacket) SetSend(value bool) {
	d.Send = value
}

type UpdatePuzzlePacket struct {
	ID        int64
	Send      bool
	Variable0 int32
}

func (packet *UpdatePuzzlePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadInt(b.Bytes(), b.Index())
	// TODO puzzle data
}

func (packet *UpdatePuzzlePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable0, b.Index())
	// TODO puzzle data
}
