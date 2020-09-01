package DSIncoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type ActionMessagePacket struct {
	*packets.DefaultPacket
	Variable0  byte
	Variable1  int32
	Variable2  int32
	Variable3  string
	Variable4  buffer.Vector2
	Color      string
	Variable6  byte
	Variable7  byte
	Variable8  bool
	Variable9  byte
	Variable10 byte
}

func init() {
	packets.ServerPackets = append(packets.ServerPackets, &ActionMessagePacket{
		DefaultPacket: new(packets.DefaultPacket),
	})
}

func (packet *ActionMessagePacket) Read(b buffer.PacketBuffer) {
}

func (packet *ActionMessagePacket) Write(b buffer.PacketBuffer) {
}
