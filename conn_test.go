package wave

import (
	"encoding/binary"
	"fmt"
	"github.com/Dmitriy-Vas/wave_buffer"
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
		conn.Close(fmt.Errorf("%s", "Hello World"))
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
	t.Data = b.ReadBytes(b.Bytes(), b.Index(), uint64(t.DataLength))
}

func (t *TestPacket) Write(b buffer.PacketBuffer) {
	b.WriteInt(b.Bytes(), int32(len(t.Data)), b.Index())
	b.WriteBytes(b.Bytes(), t.Data, b.Index())
}

func TestConn_SendPacket(t *testing.T) {
	buf.SetMaxLength(1024)
	buf.SetInitLength(8)
	connection := &Conn{
		Done:   make(chan error),
		proxy:  proxy,
		buffer: buf.Clone(),
	}

	serveConn := func(conn net.Conn, c *Conn) {
		defer conn.Close()
		buf2 := c.buffer.Clone()
		buf2.Reset()
		_, err := conn.Read(buf2.Bytes()[:8])
		assert.NoError(t, err)
		length := buf2.ReadLong(buf2.Bytes(), buf2.Index())
		buf2.Resize(uint64(length - 8))

		_, err = conn.Read(buf2.Bytes()[buf2.Index():buf2.Len()])
		assert.NoError(t, err)
		id := buf2.ReadLong(buf2.Bytes(), buf2.Index())
		assert.EqualValues(t, 15, id)
		c.Close(nil)
	}

	go func(c *Conn) {
		conn, err := listenConnection(":9124")
		assert.NoError(t, err)
		serveConn(conn, c)
	}(connection)

	go func(c *Conn) {
		conn, err := listenConnection(":9125")
		assert.NoError(t, err)
		serveConn(conn, c)
	}(connection)
	locConn, err := connectTo(":9124")
	assert.NoError(t, err)
	defer locConn.Close()
	connection.LocalConn = locConn

	remConn, err := connectTo(":9125")
	assert.NoError(t, err)
	defer remConn.Close()
	connection.RemoteConn = remConn

	testPacket := &TestPacket{
		ID:         15,
		Send:       true,
		DataLength: int32(len("Hello World")),
		Data:       []byte("Hello World"),
	}
	err = connection.SendPacket(connection.buffer, testPacket, false)
	assert.NoError(t, err)

	err = connection.SendPacket(connection.buffer, testPacket, true)
	<-connection.Done
}

func TestConn_start(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		_, err := listenConnection(":9123")
		assert.NoError(t, err)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		conn, err := listenConnection(":9124")
		assert.NoError(t, err)
		wg.Done()
		_, err = conn.Read(make([]byte, 32))
		assert.EqualError(t, err, io.EOF.Error())
		wg.Done()
	}(wg)
	proxy.Config.RemoteAddress, _ = net.ResolveTCPAddr("tcp", ":9123")
	proxy.Config.LocalAddress, _ = net.ResolveTCPAddr("tcp", ":9124")
	conn := &Conn{
		ID:    0,
		Done:  make(chan error),
		proxy: proxy,
	}
	localConn, err := connectTo(":9124")
	assert.NoError(t, err)
	conn.LocalConn = localConn
	go conn.start()
	wg.Wait()

	wg.Add(1)
	conn.Close(io.EOF)
	wg.Wait()
}

func TestConn_readAndWritePacket(t *testing.T) {
	buf := buf.Clone().(*buffer.DefaultBuffer)
	buf.SetInitLength(8)
	buf.SetMaxLength(256)
	buf.Reset()

	proxy.AddPacket(1, true, new(TestPacket))
	proxy.AddPacket(2, false, new(TestPacket))
	conn := proxy.NewConn(0, nil)

	testPacket := &TestPacket{
		ID:         1,
		Send:       true,
		DataLength: int32(len("Hello Packet")),
		Data:       []byte("Hello Packet"),
	}
	conn.writePacket(buf, testPacket)
	packet := conn.readPacket(buf, true)
	assert.EqualValues(t, packet.(*TestPacket), testPacket)
	buf.Reset()

	testPacket = &TestPacket{
		ID:         2,
		Send:       true,
		DataLength: int32(len("Hello World")),
		Data:       []byte("Hello World"),
	}
	conn.writePacket(buf, testPacket)
	packet = conn.readPacket(buf, false)
	assert.EqualValues(t, packet.(*TestPacket), testPacket)
	buf.Reset()

	conn.writePacket(buf, &TestPacket{
		ID:         3,
		Send:       true,
		DataLength: 0,
		Data:       nil,
	})
	packet = conn.readPacket(buf, true)
	assert.Nil(t, packet)
}

func TestConn_writeDataWithProcess(t *testing.T) {
	buf.SetMaxLength(1024)
	buf.SetInitLength(8)
	connection := &Conn{
		Done:   make(chan error),
		proxy:  proxy,
		buffer: buf.Clone(),
	}
	connection.proxy.Config.OutgoingProcess = func(buffer buffer.PacketBuffer, start bool) {
		if start {
			return
		}
		buffer.WriteLong(buffer.Bytes(), 941, buffer.Index())
	}

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		conn, err := listenConnection(":9123")
		assert.NoError(t, err)
		b := make([]byte, 8)
		_, err = conn.Read(b)
		assert.NoError(t, err)
		result := binary.LittleEndian.Uint64(b)
		assert.EqualValues(t, 941, result)
		wg.Done()
	}(wg)
	connection.RemoteConn, _ = connectTo(":9123")
	err := connection.writeData(connection.buffer, true)
	assert.NoError(t, err)
	wg.Wait()
}

func TestConn_pipe(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	buf := buf.Clone().(*buffer.DefaultBuffer)
	buf.SetInitLength(8)
	buf.SetMaxLength(256)
	buf.Reset()
	c := &Conn{
		ID:     0,
		Done:   make(chan error),
		proxy:  proxy,
		buffer: buf,
	}

	serve := func(conn net.Conn, wg *sync.WaitGroup) {
		packet := &TestPacket{
			ID:         1,
			Send:       true,
			DataLength: int32(len("Hello World")),
			Data:       []byte("Hello World"),
		}
		buf := c.buffer.Clone()
		c.writePacket(buf, packet)
		for i := 0; i <= 10; i++ {
			conn.Write(buf.Bytes()[:buf.Len()])
		}
		buf.Reset()
		packet = &TestPacket{
			ID:         2,
			Send:       true,
			DataLength: int32(len("Bye World")),
			Data:       []byte("Bye World"),
		}
		c.writePacket(buf, packet)
		for i := 0; i <= 10; i++ {
			conn.Write(buf.Bytes()[:buf.Len()])
		}
		wg.Done()
	}

	go func(wg *sync.WaitGroup) {
		conn, err := listenConnection(":9123")
		assert.NoError(t, err)
		serve(conn, wg)
	}(wg)
	go func(wg *sync.WaitGroup) {
		conn, err := listenConnection(":9124")
		assert.NoError(t, err)
		serve(conn, wg)
	}(wg)

	c.LocalConn, _ = connectTo(":9123")
	c.RemoteConn, _ = connectTo(":9124")

	counter := 0
	servePacket := func(conn *Conn, packet Packet) {
		if counter == 0 {
			packet.SetSend(false)
		}
		counter++
	}

	c.proxy.HookPacket(2, true, servePacket)
	c.proxy.HookPacket(2, false, servePacket)

	go c.pipe(true)
	go c.pipe(false)

	wg.Wait()
}
