package wave

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"net"
)

type Config struct {
	PacketLengthSize   uint64   // How much bytes will be used to read whole packet length
	PacketTypeSize     uint64   // How much bytes will be used to read packet type
	LengthIncludesSelf bool     // Packet length includes self extra bytes or not
	RemoteAddress      net.Addr // Target address, will be used to establish connection after receiving connection from LocalAddress
	LocalAddress       net.Addr // Local address, will be used to listen for new connection
	ConnectImmediately bool     // If true, will connect to the remote host right after receiving new connection from listener
	// Pointer to the implementation of PacketBuffer.
	// example: &buffer.DefaultBuffer{
	//		PacketReader: &buffer.DefaultReader{Order: binary.LittleEndian},
	//		PacketWriter: &buffer.DefaultWriter{Order: binary.LittleEndian},
	//	}
	Buffer buffer.PacketBuffer
	// Fires on every packet creation, you can assert type and use type-specific functions
	PacketInit func(packet Packet)
	// This functions fires twice, when packet fully received (start=true) and before sending to the destination (start=false)
	// Useful to cipher/decipher packets
	OutgoingProcess Process
	IncomingProcess Process
}

type Process func(buffer buffer.PacketBuffer, start bool)
