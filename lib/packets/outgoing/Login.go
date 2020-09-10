package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *LoginPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *LoginPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *LoginPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *LoginPacket) SetSend(value bool) {
	d.Send = value
}

type LoginPacket struct {
	ID             int64
	Send           bool
	Language       byte
	ProductVersion string
	GameVersion    int32
	Variable0      byte

	Name         string
	Password     string
	IsDiscord    bool
	IsSteam      bool
	ClientRegion string
}

func (l *LoginPacket) Read(b buffer.PacketBuffer) {
	l.Language = b.ReadByte(b.Bytes(), b.Index())
	l.ProductVersion = b.ReadString(b.Bytes(), b.Index(), 0)
	l.GameVersion = b.ReadInt(b.Bytes(), b.Index())
	l.Variable0 = b.ReadByte(b.Bytes(), b.Index())
	l.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	l.Password = b.ReadString(b.Bytes(), b.Index(), 0)
	l.IsDiscord = b.ReadBool(b.Bytes(), b.Index())
	l.IsSteam = b.ReadBool(b.Bytes(), b.Index())
	l.ClientRegion = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (l *LoginPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), l.Language, b.Index())
	b.WriteString(b.Bytes(), l.ProductVersion, b.Index())
	b.WriteInt(b.Bytes(), l.GameVersion, b.Index())
	b.WriteByte(b.Bytes(), l.Variable0, b.Index())
	b.WriteString(b.Bytes(), l.Name, b.Index())
	b.WriteString(b.Bytes(), l.Password, b.Index())
	b.WriteBool(b.Bytes(), l.IsDiscord, b.Index())
	b.WriteBool(b.Bytes(), l.IsSteam, b.Index())
	b.WriteString(b.Bytes(), l.ClientRegion, b.Index())
}
