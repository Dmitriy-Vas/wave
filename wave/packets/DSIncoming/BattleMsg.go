package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type BattleMsgPacket struct {
	*packets.DefaultPacket
	Language  int32
	Str       string
	Color     string // RGBA
	Messages  []string
	Variable6 byte
}

const (
	CustomLang = -1
)

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &BattleMsgPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *BattleMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Language = b.ReadInt(b.Bytes(), b.Index())
	packet.Str = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0) // TODO convert to RGBA
	if packet.Language != CustomLang {
		if messageAmounts := b.ReadInt(b.Bytes(), b.Index()); messageAmounts > -1 {
			packet.Messages = make([]string, messageAmounts+1)
			for i := 0; int32(i) <= messageAmounts; i++ {
				packet.Messages[i] = b.ReadString(b.Bytes(), b.Index(), 0)
			}
		}
		packet.Variable6 = b.ReadByte(b.Bytes(), b.Index())
	}
}

func (packet *BattleMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Language, b.Index())
	b.WriteString(b.Bytes(), packet.Str, b.Index())
	b.WriteString(b.Bytes(), packet.Color, b.Index())
	if packet.Language != CustomLang {
		b.WriteInt(b.Bytes(), int32(len(packet.Messages)-1), b.Index())
		for _, message := range packet.Messages {
			b.WriteString(b.Bytes(), message, b.Index())
		}
		packet.Variable6 = b.ReadByte(b.Bytes(), b.Index())
	}
}
