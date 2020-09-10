package outgoing

import (
	"github.com/Dmitriy-Vas/wave/buffer"
)

// GetID returns packet ID.
func (d *ReportPlayerPacket) GetID() int64 {
	return d.ID
}

// SetID sets ID to the packet.
func (d *ReportPlayerPacket) SetID(id int64) {
	d.ID = id
}

// GetSend returns whether to send this packet.
func (d *ReportPlayerPacket) GetSend() bool {
	return d.Send
}

// SetSend sets whether to send this packet.
func (d *ReportPlayerPacket) SetSend(value bool) {
	d.Send = value
}

type ReportPlayerPacket struct {
	ID   int64
	Send bool
	Name string
	Text string
}

func (packet *ReportPlayerPacket) Read(b buffer.PacketBuffer) {
	packet.Name = b.ReadString(b.Bytes(), b.Index(), 0)
	packet.Text = b.ReadString(b.Bytes(), b.Index(), 0)

}

func (packet *ReportPlayerPacket) Write(b buffer.PacketBuffer) {
	b.WriteString(b.Bytes(), packet.Name, b.Index())
	b.WriteString(b.Bytes(), packet.Text, b.Index())
}
