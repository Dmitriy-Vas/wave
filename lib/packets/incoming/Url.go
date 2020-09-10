package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UrlPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UrlPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UrlPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UrlPacket) SetSend(value bool) {
	d.Send = value
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
