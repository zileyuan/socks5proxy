package main

import "net"
import "github.com/zileyuan/gosocksv5d"

func main() {
	server := gosocksv5d.NewServer()
	server.ListenAndServe(net.IPv4zero, 1080) // Never returns
}
