package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *SolarEXPPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SolarEXPPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SolarEXPPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SolarEXPPacket) SetSend(value bool) {
	d.Send = value
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
