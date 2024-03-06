package main

type zSkipListLevel struct {
	forWord zSkipListNode
	span    int
}

type zSkipListNode struct {
	nodeVal  string
	score    float32
	backWord *zSkipListNode
	level    []zSkipListLevel
}

type zSkipList struct {
	headNode zSkipListNode
	tailNode zSkipListNode
	length   uint64
	level    int
}

// SkipListNode

func main() {

}

func zCreateList() {

}
