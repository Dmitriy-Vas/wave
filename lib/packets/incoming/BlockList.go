package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *BlockListPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *BlockListPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *BlockListPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *BlockListPacket) SetSend(value bool) {
	d.Send = value
}

type BlockListPacket struct {
	ID          int64
	Send        bool
	BlockedList string
	Type        byte
	Name        string
}

func (packet *BlockListPacket) Read(b buffer.PacketBuffer) {
	packet.BlockedList = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Type = b.ReadByte(b.Bytes(), b.Index())
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *BlockListPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.BlockedList, b.Index())
	b.WriteByte(b.Bytes(), packet.Type, b.Index())
	b.WriteString(b.Bytes(), packet.Name, b.Index())
}
