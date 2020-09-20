package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (packet *SpellsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *SpellsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *SpellsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *SpellsPacket) SetSend(value bool) {
	packet.Send = value
}

type SpellsPacket struct {
	ID           int64
	Send         bool
	Variable1    bool
	SpellNum     int32
	PlayerSpells []lib.PlayerSpellRec
}

func (packet *SpellsPacket) Read(b buffer.PacketBuffer) {
	packet.Variable1 = b.ReadBool(b.Bytes(), b.Index())
	packet.SpellNum = b.ReadInt(b.Bytes(), b.Index())
	packet.PlayerSpells = make([]lib.PlayerSpellRec, 42) // TODO move to constants
	for i := range packet.PlayerSpells {
		packet.PlayerSpells[i] = lib.PlayerSpellRec{
			Spell:  b.ReadInt(b.Bytes(), b.Index()),
			Uses:   int32(b.ReadLong(b.Bytes(), b.Index())),
			Master: b.ReadBool(b.Bytes(), b.Index()),
		}
	}
}

func (packet *SpellsPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.Variable1, b.Index())
	b.WriteInt(b.Bytes(), packet.SpellNum, b.Index())
	for _, spell := range packet.PlayerSpells {
		b.WriteInt(b.Bytes(), spell.Spell, b.Index())
		b.WriteLong(b.Bytes(), int64(spell.Uses), b.Index())
		b.WriteBool(b.Bytes(), spell.Master, b.Index())
	}
}
