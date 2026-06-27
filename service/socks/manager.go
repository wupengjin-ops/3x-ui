
package socks

var Pool = map[string]*Node{}

func Add(n *Node){
	Pool[n.Tag]=n
}

func List() []*Node {
	var r []*Node
	for _,v := range Pool{
		r = append(r,v)
	}
	return r
}
