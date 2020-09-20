package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *LoginOkPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *LoginOkPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *LoginOkPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *LoginOkPacket) SetSend(value bool) {
	packet.Send = value
}

type LoginOkPacket struct {
	ID      int64
	Send    bool
	MyIndex int32
	Class   int32
}

func (packet *LoginOkPacket) Read(b buffer.PacketBuffer) {
	packet.MyIndex = b.ReadInt(b.Bytes(), b.Index())
	packet.Class = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *LoginOkPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MyIndex, b.Index())
	b.WriteInt(b.Bytes(), packet.Class, b.Index())
}
