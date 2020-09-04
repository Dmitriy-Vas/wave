package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type NewAccountPacket struct {
	*wave.DefaultPacket
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
