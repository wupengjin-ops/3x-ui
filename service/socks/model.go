
package socks

type Node struct {
	Tag string
	Address string
	Port int
	Alive bool
	Latency int64
}
