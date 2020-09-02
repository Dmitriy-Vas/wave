package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type BlockListPacket struct {
	*wave.DefaultPacket
	BlockedList string
	BlockAction byte
	Variable3   string
}

const (
	UNBLOCKED byte = 0
	BLOCKED   byte = 1
	SHOW_LIST byte = 2
)

func (packet *BlockListPacket) Read(b buffer.PacketBuffer) {
	packet.BlockedList = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.BlockAction = b.ReadByte(b.Bytes(), b.Index())
	packet.Variable3 = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *BlockListPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.BlockedList, b.Index())
	b.WriteByte(b.Bytes(), packet.BlockAction, b.Index())
	b.WriteString(b.Bytes(), packet.Variable3, b.Index())
}