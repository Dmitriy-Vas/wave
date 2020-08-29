package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type ReceiveBubblePacket struct {
	*packets.DefaultPacket
	Variable0 int64
	Variable1 int32
	Variable2 int32
	Variable3 bool
	Variable4 string
	Variable5 int32
	Variable6 string
	Variable7 buffer.Vector2
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &ReceiveBubblePacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *ReceiveBubblePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadBool(b.Bytes(), b.Index())
	packet.Variable4 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable5 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable6 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable7 = b.ReadVector2(b.Bytes(), b.Index())
}

func (packet *ReceiveBubblePacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteBool(b.Bytes(), packet.Variable3, b.Index())
	b.WriteString(b.Bytes(), packet.Variable4, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable5, b.Index())
	b.WriteString(b.Bytes(), packet.Variable6, b.Index())
	b.WriteVector2(b.Bytes(), packet.Variable7, b.Index())
}
