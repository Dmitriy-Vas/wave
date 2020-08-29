package DSOutgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/packets"
)

type ClientRevisionPacket struct {
	*packets.DefaultPacket
	Unknown1 byte
	IsSteam  bool
}

func (c *ClientRevisionPacket) Read(b buffer.PacketBuffer) {
	c.Unknown1 = b.ReadByte(b.Bytes(), b.Index())
	c.IsSteam = b.ReadBool(b.Bytes(), b.Index())
}

func (c *ClientRevisionPacket) Write(b buffer.PacketBuffer) {
	b.WriteByte(b.Bytes(), c.Unknown1, b.Index())
	b.WriteBool(b.Bytes(), c.IsSteam, b.Index())
}
