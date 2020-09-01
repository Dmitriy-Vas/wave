package main

import (
	"github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/DSWrapper"
	"github.com/Dmitriy-Vas/wave/packets/DSOutgoing"
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
		BufferType:         (*DSWrapper.Buffer)(nil),
		ReaderType:         (*DSWrapper.Reader)(nil),
		WriterType:         (*DSWrapper.Writer)(nil),
		PacketInit:         wave.InitPacket,
		BufferInit:         DSWrapper.InitBuffer,
		ReaderInit:         DSWrapper.InitReader,
		WriterInit:         DSWrapper.InitWriter,
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
	proxy.AddPacket(20, true, new(DSOutgoing.ClientRevisionPacket))
	proxy.AddPacket(2, true, new(DSOutgoing.LoginPacket))
}

func RegisterHooks(proxy *wave.Proxy) {
	proxy.HookPacket(20, true, func(conn *wave.Conn, packet wave.Packet) {
		p := packet.(*DSOutgoing.ClientRevisionPacket)
		log.Printf("%+v", *p)
	})
}
