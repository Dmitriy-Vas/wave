package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *MapDataPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *MapDataPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *MapDataPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *MapDataPacket) SetSend(value bool) {
	packet.Send = value
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

func (packet *MapDataPacket) Write(_ buffer.PacketBuffer) {
}
