package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *PlayerInfoPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *PlayerInfoPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *PlayerInfoPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *PlayerInfoPacket) SetSend(value bool) {
	d.Send = value
}

type PlayerInfoPacket struct {
	ID        int64
	Send      bool
	Index     int32
	Variable2 []int32
	Variable3 []string
	ShowStats bool
}

func (packet *PlayerInfoPacket) Read(b buffer.PacketBuffer) {
	packet.Index = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable2 = make([]int32, 7) // TODO int to const
	packet.Variable3 = make([]string, 7)
	for i, _ := range packet.Variable2 {
		packet.Variable2[i] = b.ReadInt(b.Bytes(), b.Index())
		packet.Variable3[i] = b.ReadString(b.Bytes(), b.Index(), 0)
	}
	packet.ShowStats = b.ReadBool(b.Bytes(), b.Index())
}

func (packet *PlayerInfoPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Index, b.Index())
	for i, _ := range packet.Variable2 {
		b.WriteInt(b.Bytes(), packet.Variable2[i], b.Index())
		b.WriteString(b.Bytes(), packet.Variable3[i], b.Index())
	}
	b.WriteBool(b.Bytes(), packet.ShowStats, b.Index())
}
