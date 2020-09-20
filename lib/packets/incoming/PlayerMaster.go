package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *PlayerMasterPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *PlayerMasterPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *PlayerMasterPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *PlayerMasterPacket) SetSend(value bool) {
	packet.Send = value
}

type PlayerMasterPacket struct {
	ID     int64
	Send   bool
	Item   int32
	Master []int32
}

func (packet *PlayerMasterPacket) Read(b buffer.PacketBuffer) {
	packet.Item = b.ReadInt(b.Bytes(), b.Index())
	packet.Master = make([]int32, 25)
	for i := range packet.Master {
		packet.Master[i] = b.ReadInt(b.Bytes(), b.Index())
	}
}

func (packet *PlayerMasterPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Item, b.Index())
	for _, m := range packet.Master {
		b.WriteInt(b.Bytes(), m, b.Index())
	}
}
