package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *BanditPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *BanditPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *BanditPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *BanditPacket) SetSend(value bool) {
	packet.Send = value
}

type BanditPacket struct {
	ID           int64
	Send         bool
	BanditSelect int64
}

func (packet *BanditPacket) Read(b buffer.PacketBuffer) {
	packet.BanditSelect = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *BanditPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.BanditSelect, b.Index())
}
