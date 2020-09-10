package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *OpenUIPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *OpenUIPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *OpenUIPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *OpenUIPacket) SetSend(value bool) {
	d.Send = value
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
