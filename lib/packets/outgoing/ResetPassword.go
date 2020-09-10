package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ResetPasswordPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ResetPasswordPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ResetPasswordPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ResetPasswordPacket) SetSend(value bool) {
	d.Send = value
}

type ResetPasswordPacket struct {
	ID    int64
	Send  bool
	Email string
}

func (packet *ResetPasswordPacket) Read(b buffer.PacketBuffer) {
	// TODO client info data
	packet.Email = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *ResetPasswordPacket) Write(b buffer.PacketBuffer) {
	// TODO client info data
	b.WriteString(b.Bytes(), packet.Email, b.Index())
}
