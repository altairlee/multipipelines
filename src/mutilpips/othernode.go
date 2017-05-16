package pipelines

import (
	"log"
	"strconv"
	"time"
)

type preNode struct {
	node Node
	note string
}

func (p *preNode) produceData() {
	for i := 0; i < 20; i++ {
		s := "produce data : " + strconv.Itoa(i)
		log.Println(s)
		p.node.output <- s
		time.Sleep(time.Second)
	}
}
func startProduceData() *Node {
	pre := &preNode{
		note: "just do nothing",
	}
	go pre.produceData()
	return &pre.node
}

type afterNode struct {
	node   Node
	messge string
}

func (a *afterNode) processResult() {
	for {
		s := <-a.node.input
		log.Println("get final data : ", s)
		time.Sleep(time.Second)
	}
}
func startProcessData() *Node {
	after := afterNode{
		messge: "just do nothing",
	}
	go after.processResult()
	return &after.node
}