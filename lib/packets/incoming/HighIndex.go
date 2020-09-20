package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *HighIndexPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *HighIndexPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *HighIndexPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *HighIndexPacket) SetSend(value bool) {
	packet.Send = value
}

type HighIndexPacket struct {
	ID        int64
	Send      bool
	HighIndex int64
}

func (packet *HighIndexPacket) Read(b buffer.PacketBuffer) {
	packet.HighIndex = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *HighIndexPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.HighIndex, b.Index())
}
