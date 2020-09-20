package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (packet *BattleMsgPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *BattleMsgPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *BattleMsgPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *BattleMsgPacket) SetSend(value bool) {
	packet.Send = value
}

type BattleMsgPacket struct {
	ID        int64
	Send      bool
	Language  int32
	Str       string
	Color     string
	Messages  []string
	Variable6 byte
}

func (packet *BattleMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Language = b.ReadInt(b.Bytes(), b.Index())
	packet.Str = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0)
	if packet.Language != int32(lib.LangCustom) {
		if messageAmounts := b.ReadInt(b.Bytes(), b.Index()); messageAmounts > -1 {
			packet.Messages = make([]string, messageAmounts+1)
			for i := range packet.Messages {
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
	if packet.Language != int32(lib.LangCustom) {
		b.WriteInt(b.Bytes(), int32(len(packet.Messages)-1), b.Index())
		for _, message := range packet.Messages {
			b.WriteString(b.Bytes(), message, b.Index())
		}
		packet.Variable6 = b.ReadByte(b.Bytes(), b.Index())
	}
}
