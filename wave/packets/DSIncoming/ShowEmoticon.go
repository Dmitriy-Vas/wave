package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type ShowEmoticonPacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 bool
	Variable2
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &ShowEmoticonPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
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
