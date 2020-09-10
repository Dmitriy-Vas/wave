package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *WarpToMePacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *WarpToMePacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *WarpToMePacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *WarpToMePacket) SetSend(value bool) {
	d.Send = value
}

type WarpToMePacket struct {
	ID   int64
	Send bool
	Name string
}

func (packet *WarpToMePacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *WarpToMePacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
}
