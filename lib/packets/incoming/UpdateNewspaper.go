package incoming

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/buffer"
)

type UpdateNewspaperPacket struct {
	*wave.DefaultPacket
}

func (packet *UpdateNewspaperPacket) Read(b buffer.PacketBuffer) {
	// TODO Newspaper data
}

func (packet *UpdateNewspaperPacket) Write(b buffer.PacketBuffer) {
	// TODO Newspaper data
}
