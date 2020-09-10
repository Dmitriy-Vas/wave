package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *TradeStatusPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *TradeStatusPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *TradeStatusPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *TradeStatusPacket) SetSend(value bool) {
	d.Send = value
}

type TradeStatusPacket struct {
	ID        int64
	Send      bool
	Text      string
	Variable2 byte
}

func (packet *TradeStatusPacket) Read(b buffer.PacketBuffer) {

	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *TradeStatusPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
}
