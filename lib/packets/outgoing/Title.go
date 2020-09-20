package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *TitlePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *TitlePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *TitlePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *TitlePacket) SetSend(value bool) {
	packet.Send = value
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
