package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *StopDefencePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *StopDefencePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *StopDefencePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *StopDefencePacket) SetSend(value bool) {
	packet.Send = value
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
