package incoming

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
	ID     int64
	Send   bool
	Num    int32
	Prompt string
}

func (packet *BreakPointPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.Prompt = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *BreakPointPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteString(b.Bytes(), packet.Prompt, b.Index())
}
