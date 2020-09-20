package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *BreakPointPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *BreakPointPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *BreakPointPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *BreakPointPacket) SetSend(value bool) {
	packet.Send = value
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
