package main

import (
	proxy "github.com/Dmitriy-Vas/wave"
	"github.com/Dmitriy-Vas/wave/wrapper"
	"net"
)

const (
	LengthSize = 8
)

func main() {
	localAddr, _ := net.ResolveTCPAddr("tcp", ":7999")
	remoteAddr, _ := net.ResolveTCPAddr("tcp", "74.91.123.86:7000")
	P := proxy.New(remoteAddr)

	P.SetBuffer((*wrapper.Buffer)(nil), []interface{}{uint32(LengthSize)})
	P.SetReader((*wrapper.Reader)(nil))
	P.SetWriter((*wrapper.Writer)(nil))

	if err := P.Start(localAddr); err != nil {
		panic(err)
	}
	P.Close()
}
