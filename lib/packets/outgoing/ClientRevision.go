package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ClientRevisionPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ClientRevisionPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ClientRevisionPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ClientRevisionPacket) SetSend(value bool) {
	packet.Send = value
}

type ClientRevisionPacket struct {
	ID       int64
	Send     bool
	Unknown1 byte
	IsSteam  bool
}

func (packet *ClientRevisionPacket) Read(b buffer.PacketBuffer) {
	packet.Unknown1 = b.ReadByte(b.Bytes(), b.Index())
	packet.IsSteam = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *ClientRevisionPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Unknown1, b.Index())
	b.WriteBool(b.Bytes(), packet.IsSteam, b.Index())
}
