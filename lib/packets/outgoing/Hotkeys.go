package outgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type HotkeysPacket struct {
	*wave.DefaultPacket
	Hotkeys []int32
}

func (packet *HotkeysPacket) Read(b buffer.PacketBuffer) {
	packet.Hotkeys = make([]int32, 60) // TODO int to const
	for i, _ := range packet.Hotkeys {
		packet.Hotkeys[i] = b.ReadInt(b.Bytes(), b.Index())
	}

}

func (packet *HotkeysPacket) Write(b buffer.PacketBuffer) {
	for _, hotkey := range packet.Hotkeys {
		b.WriteInt(b.Bytes(), hotkey, b.Index())
	}
}
