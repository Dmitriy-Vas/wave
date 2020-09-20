package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *RequestUseCharPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *RequestUseCharPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *RequestUseCharPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *RequestUseCharPacket) SetSend(value bool) {
	packet.Send = value
}

type RequestUseCharPacket struct {
	ID           int64
	Send         bool
	SelectedChar byte
	Language     string
	IsDiscord    bool
	IsSteam      bool
}

func (packet *RequestUseCharPacket) Read(b buffer.PacketBuffer) {
	packet.SelectedChar = b.ReadByte(b.Bytes(), b.Index())
	packet.Language = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.IsDiscord = b.ReadBool(b.Bytes(), b.Index())
	packet.IsSteam = b.ReadBool(b.Bytes(), b.Index())

}

func (packet *RequestUseCharPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.SelectedChar, b.Index())
	b.WriteString(b.Bytes(), packet.Language, b.Index())
	b.WriteBool(b.Bytes(), packet.IsDiscord, b.Index())
	b.WriteBool(b.Bytes(), packet.IsSteam, b.Index())
}
