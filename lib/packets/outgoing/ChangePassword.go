package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ChangePasswordPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ChangePasswordPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ChangePasswordPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ChangePasswordPacket) SetSend(value bool) {
	packet.Send = value
}

type ChangePasswordPacket struct {
	ID       int64
	Send     bool
	Username string
	PW       string
	NPW      string
}

func (packet *ChangePasswordPacket) Read(b buffer.PacketBuffer) {
	// TODO client info data
	packet.Username = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.PW = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.NPW = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *ChangePasswordPacket) Write(b buffer.PacketBuffer) {
	// TODO client info data
	b.WriteString(b.Bytes(), packet.Username, b.Index())
	b.WriteString(b.Bytes(), packet.PW, b.Index())
	b.WriteString(b.Bytes(), packet.NPW, b.Index())
}
