package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type BlockListPacket struct {
	*wave.DefaultPacket
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
