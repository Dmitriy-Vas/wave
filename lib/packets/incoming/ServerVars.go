package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *ServerVarsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *ServerVarsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *ServerVarsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *ServerVarsPacket) SetSend(value bool) {
	packet.Send = value
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
