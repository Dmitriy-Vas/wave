package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *CheckMapPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *CheckMapPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *CheckMapPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *CheckMapPacket) SetSend(value bool) {
	packet.Send = value
}

type CheckMapPacket struct {
	ID          int64
	Send        bool
	MapNum      int32
	MapRevision int32
	X           int32
	Y           int32
}

func (packet *CheckMapPacket) Read(b buffer.PacketBuffer) {
	packet.MapNum = b.ReadInt(b.Bytes(), b.Index())
	packet.MapRevision = b.ReadInt(b.Bytes(), b.Index())
	packet.X = b.ReadInt(b.Bytes(), b.Index())
	packet.Y = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *CheckMapPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MapNum, b.Index())
	b.WriteInt(b.Bytes(), packet.MapRevision, b.Index())
	b.WriteInt(b.Bytes(), packet.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Y, b.Index())
}
