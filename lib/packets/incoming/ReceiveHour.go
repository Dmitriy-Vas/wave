package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ReceiveHourPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ReceiveHourPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ReceiveHourPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ReceiveHourPacket) SetSend(value bool) {
	d.Send = value
}

type ReceiveHourPacket struct {
	ID          int64
	Send        bool
	Hour        int32
	ServerDay   byte
	ServerMonth byte
}

func (packet *ReceiveHourPacket) Read(b buffer.PacketBuffer) {
	packet.Hour = b.ReadInt(b.Bytes(), b.Index())
	packet.ServerDay = b.ReadByte(b.Bytes(), b.Index())
	packet.ServerMonth = b.ReadByte(b.Bytes(), b.Index())
}

func (packet *ReceiveHourPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Hour, b.Index())
	b.WriteByte(b.Bytes(), packet.ServerDay, b.Index())
	b.WriteByte(b.Bytes(), packet.ServerMonth, b.Index())
}
