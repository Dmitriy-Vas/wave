package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type MapDataPacket struct {
	*wave.DefaultPacket
	Variable0 bool
	Variable1 int32
}

func (packet *MapDataPacket) Read(b buffer.PacketBuffer) {
	if packet.Variable0 = b.ReadBool(b.Bytes(), b.Index()); packet.Variable0 {
		packet.Variable1 = b.ReadInt(b.Bytes(), b.Index())
		// TODO MapData
	}
	// TODO MapCacheData
}

func (packet *MapDataPacket) Write(b buffer.PacketBuffer) {
}
