package wave

import (
	"encoding/binary"
	"fmt"
	"github.com/Dmitriy-Vas/wave/buffer"
	"github.com/stretchr/testify/assert"
	"io"
	"net"
	"sync"
	"testing"
)

var buf = &buffer.DefaultBuffer{
	PacketReader: &buffer.DefaultReader{Order: binary.LittleEndian},
	PacketWriter: &buffer.DefaultWriter{Order: binary.LittleEndian},
}
var proxy = &Proxy{
	listener:            nil,
	index:               new(uint32),
	Connections:         new(sync.Map),
	OutgoingPacketHooks: new(sync.Map),
	IncomingPacketHooks: new(sync.Map),
	OutgoingPacketMap:   new(sync.Map),
	IncomingPacketMap:   new(sync.Map),
	Config:              config,
}
var config = Config{
	PacketLengthSize:   Size64Bits,
	PacketTypeSize:     Size64Bits,
	LengthIncludesSelf: false,
	RemoteAddress:      nil,
	LocalAddress:       nil,
	ConnectImmediately: true,
	Buffer:             buf.Clone(),
}

// Automatically closes listener, but doesn't care about net.Conn
func listenConnection(address string) (net.Conn, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	defer listener.Close()

	localConn, err := listener.Accept()
	if err != nil {
		return nil, err
	}
	return localConn, nil
}

// You must close net.Conn manually
func connectTo(address string) (net.Conn, error) {
	return net.Dial("tcp", address)
}

func TestConn_Close(t *testing.T) {
	done := make(chan error)
	go func(d chan error) {
		// Blocks execution until new connection
		servConn, err := listenConnection(":9123")
		assert.NoError(t, err)
		conn := Conn{
			LocalConn: servConn,
			Done:      d,
		}
		servConn.Close()
		conn.Close(fmt.Errorf("Hello World"))
	}(done)

	clientConn, err := connectTo(":9123")
	assert.NoError(t, err)
	defer clientConn.Close()

	err = <-done
	assert.Errorf(t, err, "Hello World")
}

func TestConn_Reconnect(t *testing.T) {
	done := make(chan error)
	defer close(done)
	connection := &Conn{Done: done}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(c *Conn, wg *sync.WaitGroup) {
		// Local address
		conn, err := listenConnection(":9123")
		//log.Printf("Local connection established")
		assert.NoError(t, err)
		c.LocalConn = conn
		wg.Done()

		// Wait until c.Close()
		<-c.Done
		c.LocalConn.Close()
		c.RemoteConn.Close()
		//log.Printf("Local and Remote connections are closed")
	}(connection, wg)
	// Connect to the local address
	localConn, err := connectTo(":9123")
	assert.NoError(t, err)
	defer localConn.Close()
	wg.Wait()

	wg.Add(1)
	go func(c *Conn, wg *sync.WaitGroup) {
		// Remote address1
		conn, err := listenConnection(":9124")
		assert.NoError(t, err)
		//log.Printf(":9124 new connection received")
		wg.Done()
		_, err = conn.Read(make([]byte, 32))
		assert.EqualError(t, err, io.EOF.Error())
		//log.Printf(":9124 connection closed")
		wg.Done()
	}(connection, wg)
	//log.Printf("Connecting to :9124")
	remoteConn1, err := connectTo(":9124")
	assert.NoError(t, err)
	connection.RemoteConn = remoteConn1
	wg.Wait()

	wg.Add(1)
	go func(c *Conn, wg *sync.WaitGroup) {
		// Remote address2
		conn, err := listenConnection(":9125")
		assert.NoError(t, err)
		//log.Printf(":9125 new connection received")
		wg.Done()
		_, err = conn.Read(make([]byte, 32))
		assert.EqualError(t, err, io.EOF.Error())
		//log.Printf(":9125 connection closed")
		wg.Done()
	}(connection, wg)

	remoteAddr2, _ := net.ResolveTCPAddr("tcp", ":9125")
	//log.Printf("Connecting to :9125")
	connection.Reconnect(remoteAddr2)
	wg.Wait()
	wg.Add(2)
	connection.Close(nil)
	wg.Wait()
}

type TestPacket struct {
	ID         int64
	Send       bool
	DataLength int32
	Data       []byte
}

func (t *TestPacket) GetID() int64 {
	return t.ID
}

func (t *TestPacket) SetID(id int64) {
	t.ID = id
}

func (t *TestPacket) SetSend(value bool) {
	t.Send = value
}

func (t *TestPacket) GetSend() bool {
	return t.Send
}

func (t *TestPacket) Read(b buffer.PacketBuffer) {
	t.DataLength = b.ReadInt(b.Bytes(), b.Index())
	b.Next(4)
	t.Data = b.ReadBytes(b.Bytes(), b.Index(), uint64(t.DataLength))
	b.Next(uint64(t.DataLength))
}

func (t *TestPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), int32(len(t.Data)), b.Index())
	b.Next(4)
	b.WriteBytes(b.Bytes(), t.Data, b.Index())
	b.Next(uint64(len(t.Data)))
}

func TestConn_SendPacket(t *testing.T) {
	buf.SetMaxLength(1024)
	buf.SetInitLength(8)
	connection := &Conn{
		Done:   make(chan error),
		proxy:  proxy,
		buffer: buf.Clone(),
	}
	go func(c *Conn) {
		conn, err := listenConnection(":9124")
		assert.NoError(t, err)
		defer conn.Close()
		buf2 := c.buffer.Clone()
		buf2.Reset()
		_, err = conn.Read(buf2.Bytes()[0:16])
		assert.NoError(t, err)
		Length := binary.LittleEndian.Uint64(buf2.Bytes()[0:8])
		assert.NotEqualValues(t, 0, Length)
		ID := binary.LittleEndian.Uint64(buf2.Bytes()[8:16])
		assert.EqualValues(t, 15, ID)
		c.Close(nil)
	}(connection)
	conn, err := connectTo(":9124")
	assert.NoError(t, err)
	defer conn.Close()
	connection.LocalConn = conn

	err = connection.SendPacket(&TestPacket{
		ID:         15,
		Send:       true,
		DataLength: int32(len("Hello World")),
		Data:       []byte("Hello World"),
	}, false)
	assert.NoError(t, err)
	<-connection.Done
}
