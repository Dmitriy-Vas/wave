package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UrlPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UrlPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UrlPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UrlPacket) SetSend(value bool) {
	packet.Send = value
}

type UrlPacket struct {
	ID           int64
	Send         bool
	PromotionURL string
}

func (packet *UrlPacket) Read(b buffer.PacketBuffer) {
	packet.PromotionURL = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *UrlPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.PromotionURL, b.Index())
}
