package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *GetPlayerInfoPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *GetPlayerInfoPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *GetPlayerInfoPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *GetPlayerInfoPacket) SetSend(value bool) {
	d.Send = value
}

type GetPlayerInfoPacket struct {
	ID    int64
	Send  bool
	Index int32
	X     byte
	Y     byte
}

func (packet *GetPlayerInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.X = b.ReadByte(b.Bytes(), b.Index())
	packet.Y = b.ReadByte(b.Bytes(), b.Index())

}

func (packet *GetPlayerInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	b.WriteByte(b.Bytes(), packet.X, b.Index())
	b.WriteByte(b.Bytes(), packet.Y, b.Index())
}
