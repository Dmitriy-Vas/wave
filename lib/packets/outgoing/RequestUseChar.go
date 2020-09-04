package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type RequestUseCharPacket struct {
	*wave.DefaultPacket
	SelectedChar byte
	Language     string
	IsDiscord    bool
	IsSteam      bool
}

func (packet *RequestUseCharPacket) Read(b buffer.PacketBuffer) {
	packet.SelectedChar = b.ReadByte(b.Bytes(), b.Index())
	packet.Language = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.IsDiscord = b.ReadBool(b.Bytes(), b.Index())
	packet.IsSteam = b.ReadBool(b.Bytes(), b.Index())

}

func (packet *RequestUseCharPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.SelectedChar, b.Index())
	b.WriteString(b.Bytes(), packet.Language, b.Index())
	b.WriteBool(b.Bytes(), packet.IsDiscord, b.Index())
	b.WriteBool(b.Bytes(), packet.IsSteam, b.Index())
}
