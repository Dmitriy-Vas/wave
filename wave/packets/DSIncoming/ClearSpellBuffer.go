package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type ClearSpellBufferPacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 byte
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &ClearSpellBufferPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *ClearSpellBufferPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ClearSpellBufferPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
}
