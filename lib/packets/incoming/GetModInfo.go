package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *GetModInfoPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *GetModInfoPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *GetModInfoPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *GetModInfoPacket) SetSend(value bool) {
	d.Send = value
}

type GetModInfoPacket struct {
	ID        int64
	Send      bool
	Variable0 int64
	Variable1 byte
	Variable2 byte
	Variable3 string
	Variable4 int32
	Variable5 string
}

func (packet *GetModInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Variable4 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable5 = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *GetModInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable1, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable2, b.Index())
	b.WriteString(b.Bytes(), packet.Variable3, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable4, b.Index())
	b.WriteString(b.Bytes(), packet.Variable5, b.Index())
}
