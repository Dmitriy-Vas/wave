package wave

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"net"
)

type Config struct {
	PacketLengthSize   uint64   // How much bytes will be used to read packet length.
	PacketTypeSize     uint64   // How much bytes will be used to read packet type.
	LengthIncludesSelf bool     // Packet length includes payload size with PacketLengthSize.
	RemoteAddress      net.Addr // Target address, will be used to establish connection.
	LocalAddress       net.Addr // Local address, will be used to listen for new connection.
	ConnectImmediately bool     // If true, will connect to the remote host right after receiving new connection from listener.
	// Pointer to the implementation of PacketBuffer.
	// example: &buffer.DefaultBuffer{
	//		PacketReader: &buffer.DefaultReader{Order: binary.LittleEndian},
	//		PacketWriter: &buffer.DefaultWriter{Order: binary.LittleEndian},
	//	}
	Buffer buffer.PacketBuffer
	// This functions fires twice, when packet fully received (start=true) and before sending to the destination (start=false)
	// Useful to cipher/decipher packets
	OutgoingProcess Process
	IncomingProcess Process
}

// Process represents function to process packet data.
// Firstly executes on fully received packet (start=true).
// Next time executes right before sending to the destination (start=false).
type Process func(buffer buffer.PacketBuffer, start bool)
