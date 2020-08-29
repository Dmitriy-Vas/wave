package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type AlertMsgPacket struct {
	*packets.DefaultPacket
	Text           string
	Notify         bool
	ErrorCode      int32
	Fatal          bool
	MessagesAmount int32
	Messages       []string
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &AlertMsgPacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *AlertMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Notify = b.ReadBool(b.Bytes(), b.Index())
	packet.ErrorCode = b.ReadInt(b.Bytes(), b.Index())
	packet.Fatal = b.ReadBool(b.Bytes(), b.Index())
	packet.MessagesAmount = b.ReadInt(b.Bytes(), b.Index())
	if packet.MessagesAmount > -1 {
		packet.Messages = make([]string, 0)
		for i := 0; int32(i) <= packet.MessagesAmount; i++ {
			str := b.ReadString(b.Bytes(), b.Index(), 0)
			packet.Messages = append(packet.Messages, str)
		}
	}
}

func (packet *AlertMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteBool(b.Bytes(), packet.Notify, b.Index())
	b.WriteInt(b.Bytes(), packet.ErrorCode, b.Index())
	b.WriteBool(b.Bytes(), packet.Fatal, b.Index())
	b.WriteInt(b.Bytes(), int32(len(packet.Messages)-1), b.Index())
	for _, str := range packet.Messages {
		b.WriteString(b.Bytes(), str, b.Index())
	}
}
