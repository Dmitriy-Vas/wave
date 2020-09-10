package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerCraftingPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerCraftingPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerCraftingPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerCraftingPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerCraftingPacket struct {
	ID      int64
	Send    bool
	Num     int32
	WithLog bool
	Num2    int32
	Amount  int32
}

func (packet *PlayerCraftingPacket) Read(b buffer.PacketBuffer) {
	packet.Num = b.ReadInt(b.Bytes(), b.Index())
	packet.WithLog = b.ReadBool(b.Bytes(), b.Index())
	packet.Num2 = b.ReadInt(b.Bytes(), b.Index())
	packet.Amount = b.ReadInt(b.Bytes(), b.Index())
}

func (packet *PlayerCraftingPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Num, b.Index())
	b.WriteBool(b.Bytes(), packet.WithLog, b.Index())
	b.WriteInt(b.Bytes(), packet.Num2, b.Index())
	b.WriteInt(b.Bytes(), packet.Amount, b.Index())
}
