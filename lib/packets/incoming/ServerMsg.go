package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ServerMsgPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ServerMsgPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ServerMsgPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ServerMsgPacket) SetSend(value bool) {
	packet.Send = value
}

type ServerMsgPacket struct {
	ID    int64
	Send  bool
	Lang  int32
	Text  string
	Color string
}

func (packet *ServerMsgPacket) Read(b buffer.PacketBuffer) {
	packet.Lang = b.ReadInt(b.Bytes(), b.Index())
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Color = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *ServerMsgPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Lang, b.Index())
	b.WriteString(b.Bytes(), packet.Text, b.Index())
	b.WriteString(b.Bytes(), packet.Color, b.Index())
}
