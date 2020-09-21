package wave

import buffer "github.com/Dmitriy-Vas/wave_buffer"

type Packet interface {
	GetID() int64
	SetID(id int64)
	SetSend(value bool)
	GetSend() bool
	Read(b buffer.PacketBuffer)
	Write(b buffer.PacketBuffer)
}
