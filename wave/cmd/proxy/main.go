package main

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/packets/outgoing"
	"github.com/Dmitriy-Vas/wave/wrapper"
	"log"
	"net"
)

func main() {
	localAddr, _ := net.ResolveTCPAddr("tcp", ":7999")
	remoteAddr, _ := net.ResolveTCPAddr("tcp", "74.91.123.86:7000")
	config := wave.Config{
		PacketLengthSize:   8,
		PacketTypeSize:     8,
		RemoteAddress:      remoteAddr,
		LocalAddress:       localAddr,
		ConnectImmediately: true,
		BufferType:         (*wrapper.Buffer)(nil),
		ReaderType:         (*wrapper.Reader)(nil),
		WriterType:         (*wrapper.Writer)(nil),
		PacketInit:         wave.InitPacket,
		BufferInit:         wrapper.InitBuffer,
		ReaderInit:         wrapper.InitReader,
		WriterInit:         wrapper.InitWriter,
	}
	proxy := wave.New(config)

	RegisterPackets(proxy)
	RegisterHooks(proxy)

	if err := proxy.Start(); err != nil {
		panic(err)
	}
	proxy.Close()
}

func RegisterPackets(proxy *wave.Proxy) {
	proxy.AddPacket(20, true, new(outgoing.ClientRevisionPacket))
	proxy.AddPacket(2, true, new(outgoing.LoginPacket))
}

func RegisterHooks(proxy *wave.Proxy) {
	proxy.HookPacket(20, true, func(conn *wave.Conn, packet wave.Packet) {
		p := packet.(*outgoing.ClientRevisionPacket)
		log.Printf("%+v", *p)
	})
}
