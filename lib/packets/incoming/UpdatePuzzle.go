package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdatePuzzlePacket struct {
	*wave.DefaultPacket
	Variable0 int32
}

func (packet *UpdatePuzzlePacket) Read(b buffer.PacketBuffer) {
	packet.Variable0 = b.ReadInt(b.Bytes(), b.Index())
	// TODO puzzle data
}

func (packet *UpdatePuzzlePacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.Variable0, b.Index())
	// TODO puzzle data
}
