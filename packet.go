package wave

type Packet interface {
	GetID() int64
	SetID(id int64)
	SetSend(value bool)
	GetSend() bool
	Read(b PacketBuffer)
	Write(b PacketBuffer)
}
