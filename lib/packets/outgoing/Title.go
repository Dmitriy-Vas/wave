package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *TitlePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *TitlePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *TitlePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *TitlePacket) SetSend(value bool) {
	d.Send = value
}

type TitlePacket struct {
	ID       int64
	Send     bool
	TitleNum int32
}

func (packet *TitlePacket) Read(b buffer.PacketBuffer) {
	packet.TitleNum = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *TitlePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.TitleNum, b.Index())
}
