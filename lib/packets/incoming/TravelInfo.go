package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *TravelInfoPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *TravelInfoPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *TravelInfoPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *TravelInfoPacket) SetSend(value bool) {
	packet.Send = value
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
