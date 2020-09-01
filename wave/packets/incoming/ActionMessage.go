package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

type ActionMessagePacket struct {
	*wave.DefaultPacket
	Variable0  byte
	Variable1  int32
	Variable2  int32
	Variable3  string
	Variable4  objects.Vector2
	Color      string
	Variable6  byte
	Variable7  byte
	Variable8  bool
	Variable9  byte
	Variable10 byte
}

func (packet *ActionMessagePacket) Read(b buffer.PacketBuffer) {
}

func (packet *ActionMessagePacket) Write(b buffer.PacketBuffer) {
}
