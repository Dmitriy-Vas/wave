package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *UpdateConditionPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UpdateConditionPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UpdateConditionPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UpdateConditionPacket) SetSend(value bool) {
	packet.Send = value
}

type UpdateConditionPacket struct {
	ID        int64
	Send      bool
	Variable0 int32
}

func (packet *UpdateConditionPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadInt(b.Bytes(), b.Index())
	if v := packet.Variable0; v != 0 || v <= 400 {
		// TODO condition data
	}
}

func (packet *UpdateConditionPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable0, b.Index())
	if v := packet.Variable0; v != 0 || v <= 400 {
		// TODO condition data
	}
}
