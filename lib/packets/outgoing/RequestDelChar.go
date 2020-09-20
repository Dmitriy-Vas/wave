package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *RequestDelCharPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *RequestDelCharPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *RequestDelCharPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *RequestDelCharPacket) SetSend(value bool) {
	packet.Send = value
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
