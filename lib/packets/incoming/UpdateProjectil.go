package incoming

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
)

// GetID returns packet ID.
func (packet *UpdateProjectilPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *UpdateProjectilPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *UpdateProjectilPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *UpdateProjectilPacket) SetSend(value bool) {
	packet.Send = value
}

type UpdateProjectilPacket struct {
	ID           int64
	Send         bool
	ProjectilNum int32
	Projectil    lib.ProjectilRec
}

func (packet *UpdateProjectilPacket) Read(b buffer.PacketBuffer) {
	if packet.ProjectilNum = b.ReadInt(b.Bytes(), b.Index()); packet.ProjectilNum != 0 && packet.ProjectilNum < 100 {
		packet.Projectil = lib.ProjectilRec{
			Name:      b.ReadString(b.Bytes(), b.Index(), 0),
			Damage:    int64(b.ReadInt(b.Bytes(), b.Index())),
			Pic:       b.ReadInt(b.Bytes(), b.Index()),
			Range:     int64(b.ReadInt(b.Bytes(), b.Index())),
			Speed:     int64(b.ReadInt(b.Bytes(), b.Index())),
			Animation: int64(b.ReadInt(b.Bytes(), b.Index())),
			Light:     b.ReadBool(b.Bytes(), b.Index()),
			Int:       make([]int32, 3), // TODO move to constants
		}
		for i := range packet.Projectil.Int {
			packet.Projectil.Int[i] = b.ReadInt(b.Bytes(), b.Index())
		}
	}
}

func (packet *UpdateProjectilPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), packet.ProjectilNum, b.Index())
	if packet.ProjectilNum != 0 && packet.ProjectilNum < 100 {
		b.WriteString(b.Bytes(), packet.Projectil.Name, b.Index())
		b.WriteInt(b.Bytes(), int32(packet.Projectil.Damage), b.Index())
		b.WriteInt(b.Bytes(), packet.Projectil.Pic, b.Index())
		b.WriteInt(b.Bytes(), int32(packet.Projectil.Range), b.Index())
		b.WriteInt(b.Bytes(), int32(packet.Projectil.Speed), b.Index())
		b.WriteInt(b.Bytes(), int32(packet.Projectil.Animation), b.Index())
		b.WriteBool(b.Bytes(), packet.Projectil.Light, b.Index())
		for _, i := range packet.Projectil.Int {
			b.WriteInt(b.Bytes(), i, b.Index())
		}
	}
}
