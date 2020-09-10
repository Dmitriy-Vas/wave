package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UIBuddiesPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UIBuddiesPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UIBuddiesPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UIBuddiesPacket) SetSend(value bool) {
	d.Send = value
}

type UIBuddiesPacket struct {
	ID        int64
	Send      bool
	Action    byte
	ValueStr  string
	ValueByte byte
}

func (packet *UIBuddiesPacket) Read(b buffer.PacketBuffer) {
	packet.Action = b.ReadByte(b.Bytes(), b.Index())
	switch packet.Action {
	case 1:
		packet.ValueStr = b.ReadString(b.Bytes(), b.Index(), 0)
	case 3:
		packet.ValueByte = b.ReadByte(b.Bytes(), b.Index())
	}

}

func (packet *UIBuddiesPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Action, b.Index())
	switch packet.Action {
	case 1:
		b.WriteString(b.Bytes(), packet.ValueStr, b.Index())
	case 3:
		b.WriteByte(b.Bytes(), packet.ValueByte, b.Index())
	}
}
