package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *HotkeysPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *HotkeysPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *HotkeysPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *HotkeysPacket) SetSend(value bool) {
	d.Send = value
}

type HotkeysPacket struct {
	ID      int64
	Send    bool
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
