package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *BossMsgPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *BossMsgPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *BossMsgPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *BossMsgPacket) SetSend(value bool) {
	d.Send = value
}

type BossMsgPacket struct {
	ID       int64
	Send     bool
	Language int32
	Message  string
	Color    string
}

func (packet *BossMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Language = b.ReadInt(b.Bytes(), b.Index())
	packet.Message = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *BossMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Language, b.Index())
	b.WriteString(b.Bytes(), packet.Message, b.Index())
	b.WriteString(b.Bytes(), packet.Color, b.Index())
}
