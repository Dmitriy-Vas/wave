package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *StunnedPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *StunnedPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *StunnedPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *StunnedPacket) SetSend(value bool) {
	d.Send = value
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
