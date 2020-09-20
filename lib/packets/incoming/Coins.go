package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *CoinsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *CoinsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *CoinsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *CoinsPacket) SetSend(value bool) {
	packet.Send = value
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
