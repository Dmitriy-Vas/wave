package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ChangePasswordPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ChangePasswordPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ChangePasswordPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ChangePasswordPacket) SetSend(value bool) {
	d.Send = value
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
