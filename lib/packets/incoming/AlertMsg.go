package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *AlertMsgPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *AlertMsgPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *AlertMsgPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *AlertMsgPacket) SetSend(value bool) {
	d.Send = value
}

type AlertMsgPacket struct {
	ID        int64
	Send      bool
	Text      string
	Notify    bool
	ErrorCode int32
	Fatal     bool
	Messages  []string
}

func (packet *AlertMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Notify = b.ReadBool(b.Bytes(), b.Index())
	packet.ErrorCode = b.ReadInt(b.Bytes(), b.Index())
	packet.Fatal = b.ReadBool(b.Bytes(), b.Index())
	if messagesAmount := b.ReadInt(b.Bytes(), b.Index()); messagesAmount > -1 {
		packet.Messages = make([]string, messagesAmount)
		for i := range packet.Messages {
			packet.Messages[i] = b.ReadString(b.Bytes(), b.Index(), 0)
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
