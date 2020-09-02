package wave

import (
	"github.com/Dmitriy-Vas/wave/buffer"
	"net"
)

type Config struct {
	PacketLengthSize   uint64 // How much bytes will be used to read whole packet length
	PacketTypeSize     uint64 // How much bytes will be used to read packet type
	LengthIncludesSelf bool   // Packet length includes self extra bytes or not

	RemoteAddress      net.Addr // Target address, will be used to establish connection after receiving connection from LocalAddress
	LocalAddress       net.Addr // Local address, will be used to listen for new connection
	ConnectImmediately bool     // If true, will connect to the remote host right after receiving new connection from listener

	// TODO maybe use Prototype instead of reflection
	BufferType buffer.PacketBuffer // Pointer to the empty implementation of PacketBuffer. ex: (*DefaultBuffer)(nil)
	ReaderType buffer.PacketReader // Pointer to the empty implementation of PacketReader. ex: (*DefaultReader)(nil)
	WriterType buffer.PacketWriter // Pointer to the empty implementation of PacketWriter. ex: (*DefaultWriter)(nil)

	PacketInit func(packet Packet)              // Fires on every packet creation, you can assert type and use type-specific functions
	BufferInit func(buffer buffer.PacketBuffer) // Fires on every buffer creation, you can assert type and use type-specific functions
	ReaderInit func(reader buffer.PacketReader) // Fires on every reader creation, you can assert type and use type-specific functions
	WriterInit func(writer buffer.PacketWriter) // Fires on every writer creation, you can assert type and use type-specific functions

	// This function fires twice, when packet fully received (start=true) and before sending to the destination (start=false)
	// Useful to cipher/decipher packets
	OutgoingProcess Process
	IncomingProcess Process
}

type Process func(buffer buffer.PacketBuffer, start bool)
