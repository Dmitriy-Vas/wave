package wave

import (
	. "github.com/Dmitriy-Vas/wave/buffer"
	. "github.com/Dmitriy-Vas/wave/packets"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

type Conn struct {
	ID                         uint32
	LocalConn, RemoteConn      net.Conn
	Done                       chan error
	_proxy                     *Proxy
	clientBuffer, serverBuffer PacketBuffer
	clientHooks, serverHooks   *sync.Map
	clientMapper, serverMapper map[int64]bool // TODO delete
}

//Start starts local listener and dials to the remote host
//Note: blocking execution
func (c *Conn) Start() {
	defer func() {
		c.LocalConn.Close()
		c._proxy.Connections.Delete(c.ID)
		for key := range c.clientMapper {
			log.Printf("Client %v", key)
		}
		for key := range c.serverMapper {
			log.Printf("Server %v", key)
		}
	}()
	c.clientMapper = make(map[int64]bool)
	c.serverMapper = make(map[int64]bool)

	c._internal()
	c.UpdateHooks()
	remote_conn, err := net.Dial("tcp", c._proxy.RemoteAddr.String())
	if err != nil {
		log.Println(err)
		return
	}
	c.RemoteConn = remote_conn
	defer c.RemoteConn.Close()

	go c.Pipe(c.LocalConn, c.RemoteConn)
	go c.Pipe(c.RemoteConn, c.LocalConn)

	if err = <-c.Done; err != nil {
		log.Printf("Error occurred: %v", err)
	}
}

func (c *Conn) _internal() {
	reader := initializeStruct(
		reflect.TypeOf(c._proxy._reader),
	).(PacketReader)
	writer := initializeStruct(
		reflect.TypeOf(c._proxy._writer),
	).(PacketWriter)
	buf_type := reflect.TypeOf(c._proxy._buffer)

	c.clientBuffer = initializeStruct(buf_type).(PacketBuffer)
	c.clientBuffer.SetReader(reader)
	c.clientBuffer.SetWriter(writer)
	c.clientBuffer.Init(c._proxy._buffer_args)

	c.serverBuffer = initializeStruct(buf_type).(PacketBuffer)
	c.serverBuffer.SetReader(reader)
	c.serverBuffer.SetWriter(writer)
	c.serverBuffer.Init(c._proxy._buffer_args)
}

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

func (c *Conn) UpdateHooks() {
	c._proxy.ServerPacketHooks.Range(func(key, value interface{}) bool {
		c.serverHooks.Store(key, value)
		return true
	})
	c._proxy.ClientPacketHooks.Range(func(key, value interface{}) bool {
		c.clientHooks.Store(key, value)
		return true
	})
}

const (
	LengthSize = 8 // Long|int64
	TypeSize   = 8 // Long|int64
)

func (c *Conn) Pipe(src, dst io.ReadWriter) {
	client := src == c.LocalConn
	var buffer PacketBuffer
	var hooks *sync.Map
	if client {
		buffer = c.clientBuffer
		hooks = c.clientHooks
	} else {
		buffer = c.serverBuffer
		hooks = c.serverHooks
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
			// Payload size received
			// Payload size is without size of packet' length, need to add LengthSize
			length := buffer.ReadLong(buffer.Bytes(), buf_index) + LengthSize
			buffer.Reindex()
			buf_len = uint32(length)
			buffer.Resize(buf_len, 0)
		}
		buf_index = buffer.Next(uint32(n))
		// Wait until full packet
		if buf_index < buf_len {
			continue
		}
		// Full packet received
		buffer.Reindex()
		packet := c.ReadPacket(buffer, client)
		if packet != nil {
			hooks, ok := hooks.Load(packet.GetID())
			if !ok {
				goto write
			}
			for _, hook := range hooks.([]PacketHook) {
				hook(c, packet)
			}

			buffer.Reindex()
			buffer.WriteLong(buffer.Bytes(), 0, buffer.Index())
			buffer.WriteLong(buffer.Bytes(), packet.GetID(), buffer.Index())
			packet.Write(buffer)
			buffer.WriteLong(buffer.Bytes(), int64(len(buffer.Bytes())-LengthSize), 0)
		}
	write:
		if _, err := dst.Write(buffer.Bytes()); err != nil {
			c.Close(err)
			return
		}
		buffer.Reset()
	}
}

func (c *Conn) ReadPacket(buffer PacketBuffer, client bool) Packet {
	buffer.Next(LengthSize)
	ID := buffer.ReadLong(buffer.Bytes(), buffer.Index())

	var packet Packet
	if client {
		c.clientMapper[ID] = true
		if p, ok := c._proxy.ClientPacketMap.Load(ID); ok {
			packet = p.(Packet)
		}
	} else {
		c.serverMapper[ID] = true
		if p, ok := c._proxy.ServerPacketMap.Load(ID); ok {
			packet = p.(Packet)
		}
	}
	if packet == nil {
		return nil
	}
	packet.SetID(ID)
	packet.Read(buffer)
	return packet
}
