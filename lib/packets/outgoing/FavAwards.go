package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *FavAwardsPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *FavAwardsPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *FavAwardsPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *FavAwardsPacket) SetSend(value bool) {
	d.Send = value
}

type FavAwardsPacket struct {
	ID       int64
	Send     bool
	Action   bool
	AwardNum int32
}

func (packet *FavAwardsPacket) Read(b buffer.PacketBuffer) {
	packet.Action = b.ReadBool(b.Bytes(), b.Index())
	packet.AwardNum = b.ReadInt(b.Bytes(), b.Index())

}

func (packet *FavAwardsPacket) Write(b buffer.PacketBuffer) {
	b.WriteBool(b.Bytes(), packet.Action, b.Index())
	b.WriteInt(b.Bytes(), packet.AwardNum, b.Index())
}
