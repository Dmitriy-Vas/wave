package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *NewAccountPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *NewAccountPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *NewAccountPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *NewAccountPacket) SetSend(value bool) {
	packet.Send = value
}

type NewAccountPacket struct {
	ID             int64
	Send           bool
	Name           string
	Password       string
	RePassword     string
	Email          string
	TermsOfService bool
	ClientLanguage int32
	ClientRegion   string
}

func (packet *NewAccountPacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Password = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.RePassword = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Email = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.TermsOfService = b.ReadBool(b.Bytes(), b.Index())
	packet.ClientLanguage = b.ReadInt(b.Bytes(), b.Index())
	packet.ClientRegion = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *NewAccountPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
	b.WriteString(b.Bytes(), packet.Password, b.Index())
	b.WriteString(b.Bytes(), packet.RePassword, b.Index())
	b.WriteString(b.Bytes(), packet.Email, b.Index())
	b.WriteBool(b.Bytes(), packet.TermsOfService, b.Index())
	b.WriteInt(b.Bytes(), packet.ClientLanguage, b.Index())
	b.WriteString(b.Bytes(), packet.ClientRegion, b.Index())
}
