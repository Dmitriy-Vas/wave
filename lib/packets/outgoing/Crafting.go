package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *CraftingPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *CraftingPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *CraftingPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *CraftingPacket) SetSend(value bool) {
	packet.Send = value
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
