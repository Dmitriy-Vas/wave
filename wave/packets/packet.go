package packets

import (
	. "github.com/Dmitriy-Vas/wave/buffer"
)

type Packet interface {
	GetID() int64
	SetID(id int64)
	SetSend(value bool)
	GetSend() bool
	Read(b PacketBuffer)
	Write(b PacketBuffer)
}

type DefaultPacket struct {
	ID   int64
	Send bool
}

func (d *DefaultPacket) GetID() int64 {
	return d.ID
}

func (d *DefaultPacket) SetID(id int64) {
	d.ID = id
}

func (d *DefaultPacket) GetSend() bool {
	return d.Send
}

func (d *DefaultPacket) SetSend(value bool) {
	d.Send = value
}

var ServerPackets = make([]Packet, 0)
var ClientPackets = make([]Packet, 0)
