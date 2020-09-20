package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SolarEXPPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SolarEXPPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SolarEXPPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SolarEXPPacket) SetSend(value bool) {
	packet.Send = value
}

type SolarEXPPacket struct {
	ID     int64
	Send   bool
	Index  int64
	Amount int64
}

func (packet *SolarEXPPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadLong(b.Bytes(), b.Index())
	packet.Amount = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *SolarEXPPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Index, b.Index())
	b.WriteLong(b.Bytes(), packet.Amount, b.Index())
}
