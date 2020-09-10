package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *SayMsgPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *SayMsgPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *SayMsgPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *SayMsgPacket) SetSend(value bool) {
	d.Send = value
}

type SayMsgPacket struct {
	ID   int64
	Send bool
	Text string
}

func (packet *SayMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *SayMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Text, b.Index())
}
