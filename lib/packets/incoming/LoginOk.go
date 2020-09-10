package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *LoginOkPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *LoginOkPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *LoginOkPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *LoginOkPacket) SetSend(value bool) {
	d.Send = value
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
