package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *GetAdminInfoPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *GetAdminInfoPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *GetAdminInfoPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *GetAdminInfoPacket) SetSend(value bool) {
	d.Send = value
}

type GetAdminInfoPacket struct {
	ID        int64
	Send      bool
	Variable0 int64
	Variable1 int32
	Variable2 int32
	Variable3 int32
}

func (packet *GetAdminInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadLong(b.Bytes(), b.Index())
	packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *GetAdminInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteLong(b.Bytes(), packet.Variable0, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable2, b.Index())
	b.WriteInt(b.Bytes(), packet.Variable3, b.Index())
}
