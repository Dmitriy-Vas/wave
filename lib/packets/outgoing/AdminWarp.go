package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *AdminWarpPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *AdminWarpPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *AdminWarpPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *AdminWarpPacket) SetSend(value bool) {
	d.Send = value
}

type AdminWarpPacket struct {
	ID   int64
	Send bool
	X    int32
	Y    int32
}

func (packet *AdminWarpPacket) Read(b buffer.PacketBuffer) {
	packet.X = b.ReadInt(b.Bytes(), b.Index())
	packet.Y = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *AdminWarpPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.X, b.Index())
	b.WriteInt(b.Bytes(), packet.Y, b.Index())
}
