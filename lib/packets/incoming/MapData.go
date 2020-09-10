package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *MapDataPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *MapDataPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *MapDataPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *MapDataPacket) SetSend(value bool) {
	d.Send = value
}

type MapDataPacket struct {
	ID        int64
	Send      bool
	Variable0 bool
	Variable1 int32
}

func (packet *MapDataPacket) Read(b buffer.PacketBuffer) {
	if packet.Variable0 = b.ReadBool(b.Bytes(), b.Index()); packet.Variable0 {
		packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
		// TODO MapData
	}
	// TODO MapCacheData
}

func (packet *MapDataPacket) Write(b buffer.PacketBuffer) {
}
