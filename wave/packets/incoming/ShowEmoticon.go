package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ShowEmoticonPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 bool
	Variable2
}

func (packet *ShowEmoticonPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable2 = b.Read(b.Bytes(), b.Index())
}

func (packet *ShowEmoticonPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable1, b.Index())
	b.Write(b.Bytes(), packet.Variable2, b.Index())
}
