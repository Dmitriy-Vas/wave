package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib/wrapper"
	"time"
)

// GetID returns packet ID.
func (d *PlayerHonorPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerHonorPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerHonorPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerHonorPacket) SetSend(value bool) {
	d.Send = value
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
