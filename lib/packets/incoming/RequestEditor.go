package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *RequestEditorPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *RequestEditorPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *RequestEditorPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *RequestEditorPacket) SetSend(value bool) {
	d.Send = value
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
