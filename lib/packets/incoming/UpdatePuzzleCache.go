package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UpdatePuzzleCachePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UpdatePuzzleCachePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UpdatePuzzleCachePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UpdatePuzzleCachePacket) SetSend(value bool) {
	packet.Send = value
}

type UpdatePuzzleCachePacket struct {
	ID        int64
	Send      bool
	Variable0 int64
}

func (packet *UpdatePuzzleCachePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *UpdatePuzzleCachePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
}
