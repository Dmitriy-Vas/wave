package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type BossMsgPacket struct {
	*packets.DefaultPacket
	Language int32
	Message  string
	Color    string // RGBA
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &BossMsgPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *BossMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Language = b.ReadInt(b.Bytes(), b.Index())
	packet.Message = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0) // TODO Convert to RGBA
}

func (packet *BossMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Language, b.Index())
	b.WriteString(b.Bytes(), packet.Message, b.Index())
	b.WriteString(b.Bytes(), packet.Color, b.Index())
}
