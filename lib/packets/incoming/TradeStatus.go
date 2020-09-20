package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *TradeStatusPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *TradeStatusPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *TradeStatusPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *TradeStatusPacket) SetSend(value bool) {
	packet.Send = value
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
