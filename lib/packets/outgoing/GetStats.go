package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"time"
)

// GetID returns packet ID.
func (packet *GetStatsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *GetStatsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *GetStatsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *GetStatsPacket) SetSend(value bool) {
	packet.Send = value
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
