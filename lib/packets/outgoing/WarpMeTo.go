package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *WarpMeToPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *WarpMeToPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *WarpMeToPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *WarpMeToPacket) SetSend(value bool) {
	d.Send = value
}

type WarpMeToPacket struct {
	ID   int64
	Send bool
	Name string
}

func (packet *WarpMeToPacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *WarpMeToPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
}
