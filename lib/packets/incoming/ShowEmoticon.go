package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ShowEmoticonPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ShowEmoticonPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ShowEmoticonPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ShowEmoticonPacket) SetSend(value bool) {
	packet.Send = value
}

type ShowEmoticonPacket struct {
	ID    int64
	Send  bool
	IsPin bool
	Index int32
	Num   int32
}

func (packet *ShowEmoticonPacket) Read(b buffer.PacketBuffer) {
	packet.IsPin = b.ReadBool(b.Bytes(), b.Index())
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *ShowEmoticonPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.IsPin, b.Index())
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
}
