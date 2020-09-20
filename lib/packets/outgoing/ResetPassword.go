package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ResetPasswordPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ResetPasswordPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ResetPasswordPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ResetPasswordPacket) SetSend(value bool) {
	packet.Send = value
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
