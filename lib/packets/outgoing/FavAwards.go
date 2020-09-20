package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *FavAwardsPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *FavAwardsPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *FavAwardsPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *FavAwardsPacket) SetSend(value bool) {
	packet.Send = value
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
