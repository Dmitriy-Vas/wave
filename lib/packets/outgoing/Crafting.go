package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *CraftingPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *CraftingPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *CraftingPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *CraftingPacket) SetSend(value bool) {
	d.Send = value
}

type CraftingPacket struct {
	ID          int64
	Send        bool
	CraftingNum int32
}

func (packet *CraftingPacket) Read(b buffer.PacketBuffer) {
	packet.CraftingNum = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *CraftingPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.CraftingNum, b.Index())
}
