package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *HotkeysPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *HotkeysPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *HotkeysPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *HotkeysPacket) SetSend(value bool) {
	packet.Send = value
}

type HotkeysPacket struct {
	ID      int64
	Send    bool
	Hotkeys []int32
}

func (packet *HotkeysPacket) Read(b buffer.PacketBuffer) {
	packet.Hotkeys = make([]int32, 60) // TODO int to const
	for i := range packet.Hotkeys {
		packet.Hotkeys[i] = b.ReadInt(b.Bytes(), b.Index())
	}

}

func (packet *HotkeysPacket) Write(b buffer.PacketBuffer) {
	for _, hotkey := range packet.Hotkeys {
		b.WriteInt(b.Bytes(), hotkey, b.Index())
	}
}
