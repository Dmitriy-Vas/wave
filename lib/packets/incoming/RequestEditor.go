package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *RequestEditorPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *RequestEditorPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *RequestEditorPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *RequestEditorPacket) SetSend(value bool) {
	packet.Send = value
}

type RequestEditorPacket struct {
	ID   int64
	Send bool
	Type int32
}

func (packet *RequestEditorPacket) Read(b buffer.PacketBuffer) {
	packet.Type = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *RequestEditorPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Type, b.Index())
}
