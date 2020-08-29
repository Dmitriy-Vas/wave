package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type PlayerInvitePacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 byte
	Variable2 string
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &PlayerInvitePacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *PlayerInvitePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *PlayerInvitePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteString(b.Bytes(), packet.Variable2, b.Index())
}
