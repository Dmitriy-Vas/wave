package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type ArenaBonusPacket struct {
	*packets.DefaultPacket
	Variable1 bool
	Variable2 byte
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &ArenaBonusPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *ArenaBonusPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ArenaBonusPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
}
