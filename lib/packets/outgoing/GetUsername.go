package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type GetUsernamePacket struct {
	*wave.DefaultPacket
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
