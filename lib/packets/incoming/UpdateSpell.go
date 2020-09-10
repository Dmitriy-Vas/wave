package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *UpdateSpellPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *UpdateSpellPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *UpdateSpellPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *UpdateSpellPacket) SetSend(value bool) {
	d.Send = value
}

type UpdateSpellPacket struct {
	ID        int64
	Send      bool
	Variable0 int32
}

func (packet *UpdateSpellPacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadInt(b.Bytes(), b.Index())
	// TODO spell data
}

func (packet *UpdateSpellPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable0, b.Index())
	// TODO spell data
}
