package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PlayerTimerPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerTimerPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerTimerPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerTimerPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerTimerPacket struct {
	ID           int64
	Send         bool
	ServerTimer  int32
	TravelActive byte
}

func (packet *PlayerTimerPacket) Read(b buffer.PacketBuffer) {
	packet.ServerTimer = b.ReadInt(b.Bytes(), b.Index())
	packet.TravelActive = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *PlayerTimerPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ServerTimer, b.Index())
	b.WriteByte(b.Bytes(), packet.TravelActive, b.Index())
}
