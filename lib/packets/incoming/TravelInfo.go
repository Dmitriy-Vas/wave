package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *TravelInfoPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *TravelInfoPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *TravelInfoPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *TravelInfoPacket) SetSend(value bool) {
	d.Send = value
}

type TravelInfoPacket struct {
	ID           int64
	Send         bool
	TravelActive bool
}

func (packet *TravelInfoPacket) Read(b buffer.PacketBuffer) {
	packet.TravelActive = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *TravelInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.TravelActive, b.Index())
}
