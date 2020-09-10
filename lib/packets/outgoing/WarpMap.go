package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *WarpMapPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *WarpMapPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *WarpMapPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *WarpMapPacket) SetSend(value bool) {
	d.Send = value
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
