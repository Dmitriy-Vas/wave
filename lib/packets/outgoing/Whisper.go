package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *WhisperPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *WhisperPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *WhisperPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *WhisperPacket) SetSend(value bool) {
	d.Send = value
}

type WhisperPacket struct {
	ID   int64
	Send bool
	Name string
	Text string
}

func (packet *WhisperPacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *WhisperPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
	b.WriteString(b.Bytes(), packet.Text, b.Index())
}
