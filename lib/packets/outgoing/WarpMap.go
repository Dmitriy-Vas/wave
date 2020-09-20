package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *WarpMapPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *WarpMapPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *WarpMapPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *WarpMapPacket) SetSend(value bool) {
	packet.Send = value
}

type WarpMapPacket struct {
	ID       int64
	Send     bool
	MapIndex int32
}

func (packet *WarpMapPacket) Read(b buffer.PacketBuffer) {
	packet.MapIndex = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *WarpMapPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.MapIndex, b.Index())
}
