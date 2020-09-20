package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (packet *LoginPacket) GetID() int64 {
	return packet.ID
}

// SetID sets ID to the packet.
func (packet *LoginPacket) SetID(id int64) {
	packet.ID = id
}

// GetSend returns whether to send this packet.
func (packet *LoginPacket) GetSend() bool {
	return packet.Send
}

// SetSend sets whether to send this packet.
func (packet *LoginPacket) SetSend(value bool) {
	packet.Send = value
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

func (packet *LoginPacket) Read(b buffer.PacketBuffer) {
	packet.Language = b.ReadByte(b.Bytes(), b.Index())
	packet.ProductVersion = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.GameVersion = b.ReadInt(b.Bytes(), b.Index())
	packet.Variable0 = b.ReadByte(b.Bytes(), b.Index())
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Password = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.IsDiscord = b.ReadBool(b.Bytes(), b.Index())
	packet.IsSteam = b.ReadBool(b.Bytes(), b.Index())
	packet.ClientRegion = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (packet *LoginPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), packet.Language, b.Index())
	b.WriteString(b.Bytes(), packet.ProductVersion, b.Index())
	b.WriteInt(b.Bytes(), packet.GameVersion, b.Index())
	b.WriteByte(b.Bytes(), packet.Variable0, b.Index())
	b.WriteString(b.Bytes(), packet.Name, b.Index())
	b.WriteString(b.Bytes(), packet.Password, b.Index())
	b.WriteBool(b.Bytes(), packet.IsDiscord, b.Index())
	b.WriteBool(b.Bytes(), packet.IsSteam, b.Index())
	b.WriteString(b.Bytes(), packet.ClientRegion, b.Index())
}
