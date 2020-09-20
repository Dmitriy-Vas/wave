package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *OpenUIPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *OpenUIPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *OpenUIPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *OpenUIPacket) SetSend(value bool) {
	packet.Send = value
}

type OpenUIPacket struct {
	ID        int64
	Send      bool
	Variable1 int32
	Variable2 int32
	Item      string
}

func (packet *OpenUIPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Item = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *OpenUIPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteString(b.Bytes(), packet.Item, b.Index())
}
