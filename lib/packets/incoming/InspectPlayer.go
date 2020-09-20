package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *InspectPlayerPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *InspectPlayerPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *InspectPlayerPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *InspectPlayerPacket) SetSend(value bool) {
	packet.Send = value
}

type InspectPlayerPacket struct {
	ID        int64
	Send      bool
	Variable0 int64
	Variable1 int32
}

func (packet *InspectPlayerPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *InspectPlayerPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
}
