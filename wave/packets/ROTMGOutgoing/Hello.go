package ROTMGOutgoing

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type HelloPacket struct {
	*wave.DefaultPacket
	BuildVersion           string
	GameID                 int32
	GUID                   string
	Random1                int32
	Password               string
	Random2                int32
	Secret                 string
	KeyTime                int32
	Key                    []byte
	mapJSON                string
	EntryTag               string
	GameNet                string
	GameNetUserID          string
	PlayPlatform           string
	PlatformToken          string
	UserToken              string
	Trailer                string
	PreviousConnectionGUID string
}

func (h HelloPacket) Read(b buffer.PacketBuffer) {
	h.BuildVersion = b.ReadString(b.Bytes(), b.Index(), 0)
	h.GameID = b.ReadInt(b.Bytes(), b.Index())
	h.GUID = b.ReadString(b.Bytes(), b.Index(), 0)
	h.Random1 = b.ReadInt(b.Bytes(), b.Index())
	h.Password = b.ReadString(b.Bytes(), b.Index(), 0)
	h.Random2 = b.ReadInt(b.Bytes(), b.Index())
	h.Secret = b.ReadString(b.Bytes(), b.Index(), 0)
	h.KeyTime = b.ReadInt(b.Bytes(), b.Index())

	if length := b.ReadShort(b.Bytes(), b.Index()); length != 0 {
		h.Key = b.ReadBytes(b.Bytes(), b.Index(), uint64(length))
	}

	if length := b.ReadInt(b.Bytes(), b.Index()); length != 0 {
		h.mapJSON = string(b.ReadBytes(b.Bytes(), b.Index(), uint64(length)))
	}

	h.EntryTag = b.ReadString(b.Bytes(), b.Index(), 0)
	h.GameNet = b.ReadString(b.Bytes(), b.Index(), 0)
	h.GameNetUserID = b.ReadString(b.Bytes(), b.Index(), 0)
	h.PlayPlatform = b.ReadString(b.Bytes(), b.Index(), 0)
	h.PlatformToken = b.ReadString(b.Bytes(), b.Index(), 0)
	h.UserToken = b.ReadString(b.Bytes(), b.Index(), 0)
	h.Trailer = b.ReadString(b.Bytes(), b.Index(), 0)
	h.PreviousConnectionGUID = b.ReadString(b.Bytes(), b.Index(), 0)
}

func (h HelloPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), h.BuildVersion, b.Index())
	b.WriteInt(b.Bytes(), h.GameID, b.Index())
	b.WriteString(b.Bytes(), h.GUID, b.Index())
	b.WriteInt(b.Bytes(), h.Random1, b.Index())
	b.WriteString(b.Bytes(), h.Password, b.Index())
	b.WriteInt(b.Bytes(), h.Random2, b.Index())
	b.WriteString(b.Bytes(), h.Secret, b.Index())
	b.WriteInt(b.Bytes(), h.KeyTime, b.Index())

	b.WriteShort(b.Bytes(), int16(len(h.Key)), b.Index())
	b.WriteBytes(b.Bytes(), h.Key, b.Index())

	b.WriteInt(b.Bytes(), int32(len(h.mapJSON)), b.Index())
	b.WriteBytes(b.Bytes(), []byte(h.mapJSON), b.Index())

	b.WriteString(b.Bytes(), h.EntryTag, b.Index())
	b.WriteString(b.Bytes(), h.GameNet, b.Index())
	b.WriteString(b.Bytes(), h.GameNetUserID, b.Index())
	b.WriteString(b.Bytes(), h.PlayPlatform, b.Index())
	b.WriteString(b.Bytes(), h.PlatformToken, b.Index())
	b.WriteString(b.Bytes(), h.UserToken, b.Index())
	b.WriteString(b.Bytes(), h.Trailer, b.Index())
	b.WriteString(b.Bytes(), h.PreviousConnectionGUID, b.Index())
}
