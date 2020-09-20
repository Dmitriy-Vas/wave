package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *StunnedPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *StunnedPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *StunnedPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *StunnedPacket) SetSend(value bool) {
	packet.Send = value
}

type StunnedPacket struct {
	ID           int64
	Send         bool
	StunDuration int32
	StopDuration bool
}

func (packet *StunnedPacket) Read(b buffer.PacketBuffer) {
	packet.StunDuration = b.ReadInt(b.Bytes(), b.Index())
	packet.StopDuration = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *StunnedPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.StunDuration, b.Index())
	b.WriteBool(b.Bytes(), packet.StopDuration, b.Index())
}
