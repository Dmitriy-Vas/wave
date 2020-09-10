package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerTimerPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerTimerPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerTimerPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerTimerPacket) SetSend(value bool) {
	d.Send = value
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
