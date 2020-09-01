package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateTitlePacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 int64
	Variable2 int64
	Variable3 int64
	Variable4 int64
	Variable5 int64
}

func (packet *UpdateTitlePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadLong(b.Bytes(), b.Index())
}

func (packet *UpdateTitlePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable1, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable2, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable3, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable4, b.Index())
	b.WriteLong(b.Bytes(), packet.Variable5, b.Index())
}
