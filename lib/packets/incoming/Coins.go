package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *CoinsPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CoinsPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CoinsPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CoinsPacket) SetSend(value bool) {
	d.Send = value
}

type CoinsPacket struct {
	ID          int64
	Send        bool
	PlayerCoins int32
}

func (packet *CoinsPacket) Read(b buffer.PacketBuffer) {
	packet.PlayerCoins = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *CoinsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.PlayerCoins, b.Index())
}
