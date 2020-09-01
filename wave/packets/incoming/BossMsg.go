package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type BossMsgPacket struct {
	*wave.DefaultPacket
	Language int32
	Message  string
	Color    string // RGBA
}

func (packet *BossMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Language = b.ReadInt(b.Bytes(), b.Index())
	packet.Message = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0) // TODO Convert to RGBA
}

func (packet *BossMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Language, b.Index())
	b.WriteString(b.Bytes(), packet.Message, b.Index())
	b.WriteString(b.Bytes(), packet.Color, b.Index())
}
