package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ServerVarsPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ServerVarsPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ServerVarsPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ServerVarsPacket) SetSend(value bool) {
	d.Send = value
}

type ServerVarsPacket struct {
	ID        int64
	Send      bool
	GlobalVar int32
}

func (packet *ServerVarsPacket) Read(b buffer.PacketBuffer) {
	packet.GlobalVar = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *ServerVarsPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.GlobalVar, b.Index())
}
