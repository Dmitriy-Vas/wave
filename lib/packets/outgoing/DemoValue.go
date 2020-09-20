package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *DemoValuePacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *DemoValuePacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *DemoValuePacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *DemoValuePacket) SetSend(value bool) {
	packet.Send = value
}

type DemoValuePacket struct {
	ID      int64
	Send    bool
	ValType byte
	ValOne  int32
	ValTwo  int32
}

func (packet *DemoValuePacket) Read(b buffer.PacketBuffer) {
	packet.ValType = b.ReadByte(b.Bytes(), b.Index())
	packet.ValOne = b.ReadInt(b.Bytes(), b.Index())
	packet.ValTwo = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *DemoValuePacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.ValType, b.Index())
	b.WriteInt(b.Bytes(), packet.ValOne, b.Index())
	b.WriteInt(b.Bytes(), packet.ValTwo, b.Index())
}
