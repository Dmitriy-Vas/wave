package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *BreakPointPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *BreakPointPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *BreakPointPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *BreakPointPacket) SetSend(value bool) {
	d.Send = value
}

type BreakPointPacket struct {
	ID   int64
	Send bool
	Line int32
}

func (packet *BreakPointPacket) Read(b buffer.PacketBuffer) {
	packet.Line = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *BreakPointPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Line, b.Index())
}
