package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"time"
)

// GetID returns packet ID.
func (packet *PlayerHonorPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerHonorPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerHonorPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerHonorPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerHonorPacket struct {
	ID        int64
	Send      bool
	PlayerNum int32
	Honor     int32
	HonorDate time.Time
}

func (packet *PlayerHonorPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerNum = b.ReadInt(b.Bytes(), b.Index())
	packet.Honor = b.ReadInt(b.Bytes(), b.Index())
	packet.HonorDate = wrapper.ReadDate(b)
}

func (packet *PlayerHonorPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerNum, b.Index())
	b.WriteInt(b.Bytes(), packet.Honor, b.Index())
	wrapper.WriteDate(b, packet.HonorDate)
}
