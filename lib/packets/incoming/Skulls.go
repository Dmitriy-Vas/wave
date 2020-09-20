package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *SkullsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SkullsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SkullsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SkullsPacket) SetSend(value bool) {
	packet.Send = value
}

type SkullsPacket struct {
	ID              int64
	Send            bool
	PlayerCalaveras []int64
	PickUp          bool
	Variable3       int32
}

func (packet *SkullsPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerCalaveras = make([]int64, 2)
	for i := range packet.PlayerCalaveras {
		packet.PlayerCalaveras[i] = b.ReadLong(b.Bytes(), b.Index())
	}
	packet.PickUp = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *SkullsPacket) Write(b buffer.PacketBuffer) {
	for _, pc := range packet.PlayerCalaveras {
		b.WriteLong(b.Bytes(), pc, b.Index())
	}
	b.WriteBool(b.Bytes(), packet.PickUp, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
}
