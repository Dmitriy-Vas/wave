package wave

import (
	. "github.com/Dmitriy-Vas/wave/buffer"
	"github.com/Dmitriy-Vas/wave/lib"
	"log"
	"net"
	"reflect"
	"sync"
)

// Conn represents current connection.
type Conn struct {
	ID                    uint32   // Connection's unique ID
	LocalConn, RemoteConn net.Conn // Bidirectional streams (client <-> proxy <-> server)
	Done                  chan error

	proxy  *Proxy
	buffer PacketBuffer
}

// Start starts serving local connection.
func (c *Conn) start() {
	defer func() {
		c.LocalConn.Close()
		c.proxy.Connections.Delete(c.ID)
	}()
	c.buffer = c.proxy.Config.Buffer.Clone()

	// Some servers requires login packet in ~2 seconds after connection
	// You can disable immediate connection, hook the login packet and afterward establish connection
	if c.proxy.Config.ConnectImmediately {
		remote_conn, err := net.Dial("tcp", c.proxy.Config.RemoteAddress.String())
		if err != nil {
			log.Printf("Error connecting: %v", err)
			return
		}
		c.RemoteConn = remote_conn
		defer func() {
			if c.RemoteConn != nil {
				c.RemoteConn.Close()
			}
		}()
	}
	go c.pipe(true)
	go c.pipe(false)

	if err := <-c.Done; err != nil {
		log.Printf("Error occurred: %v", err)
	}
}

// Reconnect connects to the remote address, closes (if exists) current connection.
func (c *Conn) Reconnect(RemoteAddress net.Addr) {
	remote_conn, _ := net.Dial("tcp", RemoteAddress.String())
	if c.RemoteConn != nil {
		c.RemoteConn.Close()
	}
	c.RemoteConn = remote_conn
}

// Close closes the connection, error is optionally.
func (c *Conn) Close(err error) {
	c.Done <- err
}

// Writes buffer data to the stream.
// Executes specified Process on buffer bytes.
// Outgoing=true, if you want to write to the server.
// Outgoing=false, if you want to write to the client.
func (c *Conn) writeData(buffer PacketBuffer, outgoing bool) error {
	var conn net.Conn
	var process Process
	if outgoing {
		conn = c.RemoteConn
		process = c.proxy.Config.OutgoingProcess
	} else {
		conn = c.LocalConn
		process = c.proxy.Config.IncomingProcess
	}
	if process != nil {
		process(buffer, false)
	}
	_, err := conn.Write(buffer.Bytes()[:buffer.Len()])
	return err
}

// Writes packet data to the buffer and sends to the destination.
// Useful for manually sending packets.
// Outgoing=true, if you want to write to the server.
// Outgoing=false, if you want to write to the client.
// Concurrency UNSAFE. TODO implement concurrency safe sending packet
func (c *Conn) SendPacket(packet Packet, outgoing bool) error {
	c.buffer.Reset()
	c.writePacket(c.buffer, packet)
	return c.writeData(c.buffer, outgoing)
}

// Outgoing=true, if you want to read from client
// Outgoing=false, if you want to read from server
func (c *Conn) pipe(outgoing bool) {
	buffer := c.buffer.Clone()
	var src net.Conn
	var hooks *sync.Map
	if outgoing {
		hooks = c.proxy.OutgoingPacketHooks
	} else {
		hooks = c.proxy.IncomingPacketHooks
	}

	for {
		if outgoing {
			src = c.LocalConn
		} else {
			src = c.RemoteConn
		}
		if src == nil {
			continue
		}
		// Read bytes from the connection and saves them to the buffer
		buf_index := buffer.Index()
		buf_len := buffer.Len()
		n, err := src.Read(buffer.Bytes()[buf_index:buf_len])
		if err != nil {
			c.Close(err)
			return
		}
		if buf_index == 0 {
			// Payload size
			length := ReadNumber(buffer, c.proxy.Config.PacketLengthSize)
			if !c.proxy.Config.LengthIncludesSelf {
				length += int64(c.proxy.Config.PacketLengthSize)
			}
			// Returns back in index
			buffer.Back(c.proxy.Config.PacketLengthSize)
			buf_len = uint64(length)
			// Grow up buffer to fit full packet
			buffer.Resize(buf_len)
		}
		buf_index = buffer.Next(uint64(n))
		// Wait until full packet
		if buf_index < buf_len {
			continue
		}
		// Full packet received
		buffer.Back(buf_index - c.proxy.Config.PacketLengthSize)
		packet := c.readPacket(buffer, outgoing)
		if packet != nil {
			if outgoing {
				log.Printf("[%s]: %+v", lib.ClientPacket(packet.GetID()).String(), packet)
			} else {
				log.Printf("[%s]: %+v", lib.ServerPacket(packet.GetID()).String(), packet)
			}
			hooks, ok := hooks.Load(packet.GetID())
			if !ok {
				goto write
			}
			for _, hook := range hooks.([]PacketHook) {
				hook(c, packet)
			}
			if !packet.GetSend() {
				buffer.Reset()
				continue
			}

			buffer.Back(buffer.Index())
			c.writePacket(buffer, packet)
		}
	write:
		if err := c.writeData(buffer, outgoing); err != nil {
			c.Close(err)
			return
		}
		buffer.Reset()
	}
}

func (c *Conn) readPacket(buffer PacketBuffer, outgoing bool) Packet {
	ID := ReadNumber(buffer, c.proxy.Config.PacketTypeSize)
	var packetType reflect.Type = nil
	if outgoing {
		if p, ok := c.proxy.OutgoingPacketMap.Load(ID); ok {
			packetType = p.(reflect.Type)
		}
	} else {
		if p, ok := c.proxy.IncomingPacketMap.Load(ID); ok {
			packetType = p.(reflect.Type)
		}
	}
	if packetType == nil {
		return nil
	}
	packet := InitializeStruct(packetType).(Packet)
	packet.SetID(ID)
	packet.Read(buffer)
	return packet
}

func (c *Conn) writePacket(buffer PacketBuffer, packet Packet) {
	// Reserve space for packet length
	WriteNumber(buffer, c.proxy.Config.PacketLengthSize, 0)
	// Write packet ID
	WriteNumber(buffer, c.proxy.Config.PacketTypeSize, packet.GetID())
	// Write whole packet data (payload)
	packet.Write(buffer)

	// Returns index to the beginning
	buffer.Back(buffer.Index())
	size := buffer.Len()
	if !c.proxy.Config.LengthIncludesSelf {
		size -= c.proxy.Config.PacketLengthSize
	}
	// Write actual packet size
	WriteNumber(buffer, c.proxy.Config.PacketLengthSize, int64(size))
}
