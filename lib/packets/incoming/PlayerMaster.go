package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerMasterPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerMasterPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerMasterPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerMasterPacket) SetSend(value bool) {
	d.Send = value
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
	for i, _ := range packet.Master {
		packet.Master[i] = b.ReadInt(b.Bytes(), b.Index())
	}
}

func (packet *PlayerMasterPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Item, b.Index())
	for _, m := range packet.Master {
		b.WriteInt(b.Bytes(), m, b.Index())
	}
}
