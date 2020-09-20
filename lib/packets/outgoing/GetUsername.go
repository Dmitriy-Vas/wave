package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *GetUsernamePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *GetUsernamePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *GetUsernamePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *GetUsernamePacket) SetSend(value bool) {
	packet.Send = value
}

type GetUsernamePacket struct {
	ID    int64
	Send  bool
	Email string
}

func (packet *GetUsernamePacket) Read(b buffer.PacketBuffer) {
	// TODO client info data
	packet.Email = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *GetUsernamePacket) Write(b buffer.PacketBuffer) {
	// TODO client info data
	b.WriteString(b.Bytes(), packet.Email, b.Index())
}
