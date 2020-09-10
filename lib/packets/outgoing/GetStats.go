package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"time"
)

// GetID returns packet ID.
func (d *GetStatsPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *GetStatsPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *GetStatsPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *GetStatsPacket) SetSend(value bool) {
	d.Send = value
}

type GetStatsPacket struct {
	ID        int64
	Send      bool
	Variable1 time.Time
	Variable2 time.Time
}

func (packet *GetStatsPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = wrapper.ReadDate(b)
	packet.Variable2 = wrapper.ReadDate(b)

}

func (packet *GetStatsPacket) Write(b buffer.PacketBuffer) {
	wrapper.WriteDate(b, packet.Variable1)
	wrapper.WriteDate(b, packet.Variable2)
}
