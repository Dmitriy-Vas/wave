package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ServerStatsPacket struct {
	*wave.DefaultPacket
	Variable0 int64
	Variable1 byte
	Variable2 string
	Variable3 string
	Variable4 int32
}

func (packet *ServerStatsPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable3 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable4 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *ServerStatsPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteString(b.Bytes(), packet.Variable2, b.Index())
	b.WriteString(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable4, b.Index())
}
