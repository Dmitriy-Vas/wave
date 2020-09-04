package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ChangePasswordPacket struct {
	*wave.DefaultPacket
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
