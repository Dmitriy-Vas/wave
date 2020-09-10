package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *StopDefencePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *StopDefencePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *StopDefencePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *StopDefencePacket) SetSend(value bool) {
	d.Send = value
}

type StopDefencePacket struct {
	ID        int64
	Send      bool
	Variable1 int64
}

func (packet *StopDefencePacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadLong(b.Bytes(), b.Index())

}

func (packet *StopDefencePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable1, b.Index())
}
