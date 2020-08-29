package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type QuestMsgPacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 string
	Variable2 string
	Variable3 int32
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &QuestMsgPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *QuestMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable2 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *QuestMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteString(b.Bytes(), packet.Variable1, b.Index())
	b.WriteString(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
}
