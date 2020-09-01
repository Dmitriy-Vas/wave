package wave

import (
	. "github.com/Dmitriy-Vas/wave/buffer"
	"log"
	"net"
	"reflect"
	"sync"
)

type Conn struct {
	ID                    uint32
	LocalConn, RemoteConn net.Conn
	Done                  chan error

	proxy                      *Proxy
	clientBuffer, serverBuffer PacketBuffer
}

//Start starts local listener and dials to the remote host
func (c *Conn) Start() {
	defer func() {
		c.LocalConn.Close()
		c.proxy.Connections.Delete(c.ID)
	}()
	c.init()

	// Some servers requires sending login packet in ~2 seconds after connection
	if c.proxy.Config.ConnectImmediately {
		remote_conn, err := net.Dial("tcp", c.proxy.Config.RemoteAddress.String())
		if err != nil {
			log.Printf("Error connecting: %v", err)
			return
		}
		c.RemoteConn = remote_conn
		defer c.RemoteConn.Close()
	}
	go c.Pipe(true)
	go c.Pipe(false)

	if err := <-c.Done; err != nil {
		log.Printf("Error occurred: %v", err)
	}
}

func (c *Conn) Reconnect(RemoteAddress net.Addr) {
	remote_conn, err := net.Dial("tcp", RemoteAddress.String())
	if err != nil {
		c.Close(err)
		return
	}
	c.RemoteConn.Close()
	c.RemoteConn = remote_conn
}

// Creates new instance of specified type
func initializeStruct(t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.Ptr:
		return reflect.New(t.Elem()).Interface()
	case reflect.Struct:
		return reflect.New(t).Interface()
	}
	return nil
}

//Close closes connection, error is optionally
func (c *Conn) Close(err error) {
	c.Done <- err
}

// TODO create better way to pass packets manually
func (c *Conn) Send(buffer PacketBuffer, client bool, process Process) error {
	var conn net.Conn
	if client {
		conn = c.LocalConn
	} else {
		conn = c.RemoteConn
	}
	if process != nil {
		process(buffer, false)
	}
	_, err := conn.Write(buffer.Bytes())
	return err
}

func (c *Conn) Pipe(client bool) {
	var buffer PacketBuffer
	var hooks *sync.Map
	var src net.Conn
	var process Process
	if client {
		src = c.LocalConn
		buffer = c.clientBuffer
		hooks = c.proxy.ClientPacketHooks
		process = c.proxy.Config.OutgoingProcess
	} else {
		src = c.RemoteConn
		buffer = c.serverBuffer
		hooks = c.proxy.ServerPacketHooks
		process = c.proxy.Config.IncomingProcess
	}

	for {
		// Reads bytes from the connection and saves them to the buffer
		buf_index := buffer.Index()
		n, err := src.Read(buffer.Bytes()[buf_index:])
		if err != nil {
			c.Close(err)
			return
		}
		buf_len := buffer.Len()
		if buf_index == 0 {
			// Payload size
			length := c.ReadNumber(buffer, c.proxy.Config.PacketLengthSize)
			if !c.proxy.Config.LengthIncludesSelf {
				length += int64(c.proxy.Config.PacketLengthSize)
			}
			// Returns back in index
			buffer.Back(c.proxy.Config.PacketLengthSize)
			buf_len = uint64(length)
			buffer.Resize(buf_len)
		}
		buf_index = buffer.Next(uint64(n))
		// Wait until full packet
		if buf_index < buf_len {
			continue
		}
		// Full packet received
		buffer.Back(buf_index - c.proxy.Config.PacketLengthSize)
		packet := c.ReadPacket(buffer, client, process)
		if packet != nil {
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
			c.WritePacket(buffer, packet)
		}
	write:
		if err := c.Send(buffer, !client, process); err != nil {
			c.Close(err)
			return
		}
		buffer.Reset()
	}
}

func (c *Conn) ReadPacket(buffer PacketBuffer, client bool, process Process) Packet {
	if process != nil {
		process(buffer, true)
	}
	ID := c.ReadNumber(buffer, c.proxy.Config.PacketTypeSize)

	var packetType reflect.Type = nil
	if client {
		if p, ok := c.proxy.ClientPacketMap.Load(ID); ok {
			packetType = p.(reflect.Type)
		}
	} else {
		if p, ok := c.proxy.ServerPacketMap.Load(ID); ok {
			packetType = p.(reflect.Type)
		}
	}
	if packetType == nil {
		return nil
	}
	packet := initializeStruct(packetType).(Packet)
	if init := c.proxy.Config.PacketInit; init != nil {
		init(packet)
	}
	packet.SetID(ID)
	packet.Read(buffer)
	return packet
}

func (c *Conn) WritePacket(buffer PacketBuffer, packet Packet) {
	// Reserve space for packet length
	c.WriteNumber(buffer, c.proxy.Config.PacketLengthSize, 0)
	// Write packet ID
	c.WriteNumber(buffer, c.proxy.Config.PacketTypeSize, packet.GetID())
	// Write whole packet data (payload)
	packet.Write(buffer)
	// Write actual packet size
	buffer.Back(buffer.Index())
	size := buffer.Len()
	if !c.proxy.Config.LengthIncludesSelf {
		size -= c.proxy.Config.PacketLengthSize
	}
	c.WriteNumber(buffer, c.proxy.Config.PacketLengthSize, int64(size))
}

const (
	Size64Bits = 8
	Size32Bits = 4
	Size16Bits = 2
	Size8Bits  = 1
)

func (c *Conn) ReadNumber(buffer PacketBuffer, size uint64) (result int64) {
	switch size {
	case Size64Bits:
		result = buffer.ReadLong(buffer.Bytes(), buffer.Index())
	case Size32Bits:
		result = int64(buffer.ReadInt(buffer.Bytes(), buffer.Index()))
	case Size16Bits:
		result = int64(buffer.ReadShort(buffer.Bytes(), buffer.Index()))
	case Size8Bits:
		result = int64(buffer.ReadByte(buffer.Bytes(), buffer.Index()))
	}
	return result
}

func (c *Conn) WriteNumber(buffer PacketBuffer, size uint64, value int64) {
	switch size {
	case Size64Bits:
		buffer.WriteLong(buffer.Bytes(), value, buffer.Index())
	case Size32Bits:
		buffer.WriteInt(buffer.Bytes(), int32(value), buffer.Index())
	case Size16Bits:
		buffer.WriteShort(buffer.Bytes(), int16(value), buffer.Index())
	case Size8Bits:
		buffer.WriteByte(buffer.Bytes(), byte(value), buffer.Index())
	}
}

func (c *Conn) init() {
	bufferType := reflect.TypeOf(c.proxy.Config.BufferType)
	readerType := reflect.TypeOf(c.proxy.Config.ReaderType)
	writerType := reflect.TypeOf(c.proxy.Config.WriterType)

	clientBuffer := initializeStruct(bufferType).(PacketBuffer)
	serverBuffer := initializeStruct(bufferType).(PacketBuffer)
	reader := initializeStruct(readerType).(PacketReader)
	writer := initializeStruct(writerType).(PacketWriter)

	if init := c.proxy.Config.BufferInit; init != nil {
		init(clientBuffer)
		init(serverBuffer)
	}
	if init := c.proxy.Config.ReaderInit; init != nil {
		init(reader)
	}
	if init := c.proxy.Config.WriterInit; init != nil {
		init(writer)
	}
	clientBuffer.SetReader(reader)
	clientBuffer.SetWriter(writer)
	serverBuffer.SetReader(reader)
	serverBuffer.SetWriter(writer)

	c.clientBuffer = clientBuffer
	c.serverBuffer = serverBuffer
}
