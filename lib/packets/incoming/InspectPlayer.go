package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type InspectPlayerPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 int32
}

func (packet *InspectPlayerPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *InspectPlayerPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
}
