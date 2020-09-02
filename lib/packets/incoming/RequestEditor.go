package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type RequestEditorPacket struct {
	*wave.DefaultPacket
	Type int32
}

func (packet *RequestEditorPacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *RequestEditorPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Type, b.Index())
}
