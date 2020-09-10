package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *RequestDelCharPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *RequestDelCharPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *RequestDelCharPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *RequestDelCharPacket) SetSend(value bool) {
	d.Send = value
}

type RequestDelCharPacket struct {
	ID           int64
	Send         bool
	SelectedChar byte
}

func (packet *RequestDelCharPacket) Read(b buffer.PacketBuffer) {
	packet.SelectedChar = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *RequestDelCharPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.SelectedChar, b.Index())
}
